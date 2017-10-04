package server

import (
	"crypto/rsa"
	"flag"
	pb "github.com/PomeloCloud/BFTRaft4go/proto"
	"github.com/dgraph-io/badger"
	"github.com/patrickmn/go-cache"
	context "golang.org/x/net/context"
	"google.golang.org/grpc"
	"net"
	"sync"
	"time"
	"bytes"
)

type Options struct {
	MaxReplications  uint32
	DBPath           string
	Address          string
	PrivateKey       []byte
	ConsensusTimeout time.Duration
}

type BFTRaftServer struct {
	Id                uint64
	Opts              Options
	DB                *badger.KV
	FuncReg           map[uint64]map[uint64]func(arg []byte) []byte
	Peers             *cache.Cache
	Groups            *cache.Cache
	GroupsPeers       *cache.Cache
	Nodes             *cache.Cache
	GroupAppendedLogs *cache.Cache
	GroupAppendLogRes *cache.Cache
	PrivateKey        *rsa.PrivateKey
	Clients           ClientStore
	lock              *sync.RWMutex
}

func (s *BFTRaftServer) ExecCommand(ctx context.Context, cmd *pb.CommandRequest) (*pb.CommandResponse, error) {
	group_id := cmd.Group
	group := s.GetGroup(group_id)
	if group == nil {
		return nil, nil
	}
	leader_peer_id := group.LeaderPeer
	leader_peer := s.GetPeer(group_id, leader_peer_id)
	response := &pb.CommandResponse{
		Group:     cmd.Group,
		LeaderId:  0,
		NodeId:    s.Id,
		RequestId: cmd.RequestId,
		Signature: s.Sign(U64Bytes(cmd.RequestId)),
		Result:    []byte{},
	}
	if leader_peer == nil {
		return response, nil
	}
	if leader_node := s.GetNode(leader_peer.Host); leader_node != nil {
		if leader_node.Id != s.Id {
			if client, err := s.Clients.Get(leader_node.ServerAddr); err != nil {
				return client.rpc.ExecCommand(ctx, cmd)
			}
		} else { // the node is the leader to this group
			response.LeaderId = leader_peer.Id
			index := s.IncrGetGroupLogLastIndex(group_id)
			hash, _ := LogHash(s.LastEntryHash(group_id), index)
			logEntry := pb.LogEntry{
				Term:    group.Term,
				Index:   index,
				Hash:    hash,
				Command: cmd,
			}
			if s.AppendEntryToLocal(group, logEntry) == nil {
				s.SendFollowersHeartbeat(ctx, group)
				response.Result = *s.WaitLogAppended(group_id, index)
			}
		}
	}
	return response, nil
}

func (s *BFTRaftServer) AppendEntries(ctx context.Context, req *pb.AppendEntriesRequest) (*pb.AppendEntriesResponse, error) {
	groupId := req.Group
	group := s.GetGroup(groupId)
	reqLeaderId := req.LeaderId
	leaderPeer := s.GetPeer(groupId, reqLeaderId)
	lastLogHash := s.LastEntryHash(groupId)
	response := &pb.AppendEntriesResponse{
		Group:     groupId,
		Term:      group.Term,
		Index:     s.GetGroupLogLastIndex(groupId),
		NodeId:    s.Id,
		Successed: false,
		Convinced: false,
		Hash:      lastLogHash,
		Signature: s.Sign(lastLogHash),
	}
	// verify group and leader existence
	if group == nil || leaderPeer == nil {
		return response, nil
	}
	// check leader transfer
	if group.LeaderPeer != reqLeaderId { // TODO: check for leader transfer
		return response, nil
	}
	// check leader node exists
	leaderNode := s.GetNode(leaderPeer.Host)
	if leaderPeer == nil {
		return response, nil
	}
	// verify signature
	if leaderPublicKey, err := ParsePublicKey(leaderNode.PublicKey); err != nil {
		signData := AppendLogEntrySignData(group.Id, group.Term, req.PrevLogIndex, req.PrevLogTerm)
		if VerifySign(leaderPublicKey, req.Signature, signData) != nil {
			return response, nil
		}
	} else {
		return response, nil
	}
	if len(req.Entries) > 0 {
		// check last log matches the first provided by the leader
		// this strategy assumes split brain will never happened (on internet)
		// the leader will always provide the entries no more than it needed
		// if the leader failed to provide the right first entries, the follower
		// 	 will not commit the log but response with current log index and term instead
		// the leader should response immediately for failed follower response
		lastLogEntry := s.LastLogEntry(groupId)
		lastLogIdx := s.GetGroupLogLastIndex(groupId)
		nextLogIdx := lastLogIdx + 1
		if lastLogEntry.Index != lastLogIdx {
			// this is unexpected, should be detected quickly
			panic("Log list unmatched with last log index")
		}
		if req.PrevLogIndex == lastLogIdx && req.Entries[0].Index == nextLogIdx { // index matched
			if req.PrevLogTerm != lastLogEntry.Term {
				// log mismatch, cannot preceded
				// what to do next will leave to the leader
				return response, nil
			} else {
				// first log matched
				// but we still need to check hash for upcoming logs
				expectedHash := s.LastEntryHash(groupId)
				for i := nextLogIdx; i < nextLogIdx + uint64(len(req.Entries)); i++ {
					expectedHash, _ = LogHash(expectedHash, i)
					entry := req.Entries[i - nextLogIdx]
					if entry.Index != i || !bytes.Equal(entry.Hash, expectedHash) {
						// not all entries match
						return response, nil
					}
				}

			}
		} else {
			return response, nil
		}
	}
	return nil, nil
}

func (s *BFTRaftServer) RequestVote(context.Context, *pb.RequestVoteRequest) (*pb.RequestVoteResponse, error) {
	return nil, nil
}

func (s *BFTRaftServer) RegisterServerFunc(group uint64, func_id uint64, fn func(arg []byte) []byte) {
	s.FuncReg[group][func_id] = fn
}

func (s *BFTRaftServer) SendFollowersHeartbeat(ctx context.Context, group *pb.RaftGroup) {
	group_peers := s.GetGroupPeers(group.Id)
	host_peers := map[*pb.Peer]bool{}
	for _, peer := range group_peers {
		host_peers[peer] = true
	}
	for peer := range host_peers {
		s.SendPeerUncommittedLogEntries(ctx, group, peer)
	}
}

func start(serverOpts Options) error {
	flag.Parse()
	lis, err := net.Listen("tcp", serverOpts.Address)
	if err != nil {
		return err
	}
	dbopt := badger.DefaultOptions
	dbopt.Dir = serverOpts.DBPath
	dbopt.ValueDir = serverOpts.DBPath
	db, err := badger.NewKV(&dbopt)
	if err != nil {
		return err
	}
	grpcServer := grpc.NewServer()
	config, err := GetConfig(db)
	if err != nil {
		return err
	}
	privateKey, err := ParsePrivateKey(config.PublicKey)
	if err != nil {
		return err
	}
	bftRaftServer := BFTRaftServer{
		Id:                HashPublicKey(PublicKeyFromPrivate(privateKey)),
		Opts:              serverOpts,
		DB:                db,
		Clients:           NewClientStore(),
		Groups:            cache.New(1*time.Minute, 1*time.Minute),
		GroupsPeers:       cache.New(1*time.Minute, 1*time.Minute),
		Peers:             cache.New(1*time.Minute, 1*time.Minute),
		Nodes:             cache.New(1*time.Minute, 1*time.Minute),
		GroupAppendLogRes: cache.New(2*time.Minute, 1*time.Minute),
		GroupAppendedLogs: cache.New(5*time.Minute, 1*time.Minute),
		PrivateKey:        privateKey,
	}
	pb.RegisterBFTRaftServer(grpcServer, &bftRaftServer)
	grpcServer.Serve(lis)
	return nil
}
