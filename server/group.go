package server

import (
	pb "github.com/PomeloCloud/BFTRaft4go/proto/server"
	"github.com/dgraph-io/badger"
	"github.com/golang/protobuf/proto"
	"github.com/patrickmn/go-cache"
	"github.com/tevino/abool"
	"log"
	"math/rand"
	"strconv"
	sync "github.com/90TechSAS/go-recursive-mutex"
	"time"
)

const (
	LEADER    = 0
	FOLLOWER  = 1
	CANDIDATE = 2
	OBSERVER  = 3
)

type RTGroupMeta struct {
	Peer              uint64
	Leader            uint64
	VotedPeer         uint64
	GroupPeers        map[uint64]*pb.Peer
	Group             *pb.RaftGroup
	Timeout           time.Time
	Role              int
	Votes             []*pb.RequestVoteResponse
	SendVotesForPeers map[uint64]bool // key is peer id
	IsBusy            *abool.AtomicBool
	Lock              *sync.RecursiveMutex
}

func NewRTGroupMeta(
	peer uint64,
	leader uint64,
	groupPeers map[uint64]*pb.Peer,
	group *pb.RaftGroup,
) *RTGroupMeta {
	return &RTGroupMeta{
		Peer:              peer,
		Leader:            leader,
		VotedPeer:         0,
		GroupPeers:        groupPeers,
		Group:             group,
		Timeout:           time.Now().Add(10 * time.Second),
		Role:              FOLLOWER,
		Votes:             []*pb.RequestVoteResponse{},
		SendVotesForPeers: map[uint64]bool{},
		IsBusy:            abool.NewBool(false),
		Lock:              &sync.RecursiveMutex{},
	}
}

func GetGroupFromKV(txn *badger.Txn, groupId uint64) *pb.RaftGroup {
	group := &pb.RaftGroup{}
	keyPrefix := ComposeKeyPrefix(groupId, GROUP_META)
	if item, err := txn.Get(keyPrefix); err == nil {
		data := ItemValue(item)
		if data == nil {
			return nil
		} else {
			proto.Unmarshal(*data, group)
			return group
		}
	} else {
		return nil
	}
}

func (s *BFTRaftServer) GetGroup(txn *badger.Txn, groupId uint64) *pb.RaftGroup {
	cacheKey := strconv.Itoa(int(groupId))
	cachedGroup, cacheFound := s.Groups.Get(cacheKey)
	if cacheFound {
		return cachedGroup.(*pb.RaftGroup)
	} else {
		group := GetGroupFromKV(txn, groupId)
		if group != nil {
			s.Groups.Set(cacheKey, group, cache.DefaultExpiration)
			return group
		} else {
			return nil
		}
	}
}

func (s *BFTRaftServer) GetGroupNTXN(groupId uint64) *pb.RaftGroup {
	group := &pb.RaftGroup{}
	s.DB.View(func(txn *badger.Txn) error {
		group = s.GetGroup(txn, groupId)
		return nil
	})
	return group
}

func (s *BFTRaftServer) SaveGroup(txn *badger.Txn, group *pb.RaftGroup) error {
	if data, err := proto.Marshal(group); err == nil {
		dbKey := ComposeKeyPrefix(group.Id, GROUP_META)
		return txn.Set(dbKey, data, 0x00)
	} else {
		return err
	}
}

func (s *BFTRaftServer) SaveGroupNTXN(group *pb.RaftGroup) error {
	return s.DB.Update(func(txn *badger.Txn) error {
		return s.SaveGroup(txn, group)
	})
}

func (s *BFTRaftServer) GetGroupHosts(txn *badger.Txn, groupId uint64) []*pb.Host {
	nodes := []*pb.Host{}
	peers := GetGroupPeersFromKV(txn, groupId)
	for _, peer := range peers {
		node := s.GetHost(txn, peer.Id)
		if node != nil {
			nodes = append(nodes, node)
		} else {
			log.Println(s.Id, "cannot find group node:", peer.Id)
		}
	}
	return nodes
}

func (s *BFTRaftServer) GetGroupHostsNTXN(groupId uint64) []*pb.Host {
	result := []*pb.Host{}
	s.DB.View(func(txn *badger.Txn) error {
		result = s.GetGroupHosts(txn, groupId)
		return nil
	})
	return result
}

func (s *BFTRaftServer) GroupLeader(groupId uint64) *pb.GroupLeader {
	res := &pb.GroupLeader{}
	if meta, onboard := s.GroupsOnboard[groupId]; onboard {
		node := s.GetHostNTXN(meta.Leader)
		if node == nil {
			log.Println("cannot get node for group leader")
		}
		res = &pb.GroupLeader{
			Node:    node,
			Accuate: true,
		}
	} else {
		// group not on the host
		// will select a host randomly in the group
		res := &pb.GroupLeader{
			Accuate: false,
		}
		s.DB.View(func(txn *badger.Txn) error {
			hosts := s.GetGroupHosts(txn, groupId)
			if len(hosts) > 0 {
				res.Node = hosts[rand.Intn(len(hosts))]
			} else {
				log.Println("cannot get group leader")
			}
			return nil
		})
	}
	return res
}
