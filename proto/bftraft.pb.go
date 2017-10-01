// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/bftraft.proto

/*
Package bftraft is a generated protocol buffer package.

It is generated from these files:
	proto/bftraft.proto

It has these top-level messages:
	CommandRequest
	CommandResponse
	LogEntry
	RequestVoteRequest
	RequestVoteResponse
	AppendEntriesRequest
	AppendEntriesResponse
	Peer
	Node
	RaftGroup
*/
package bftraft

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type CommandRequest struct {
	Group     uint64 `protobuf:"varint,1,opt,name=group" json:"group,omitempty"`
	ClientId  uint64 `protobuf:"varint,2,opt,name=client_id,json=clientId" json:"client_id,omitempty"`
	RequestId uint64 `protobuf:"varint,3,opt,name=request_id,json=requestId" json:"request_id,omitempty"`
	FuncId    uint64 `protobuf:"varint,4,opt,name=func_id,json=funcId" json:"func_id,omitempty"`
	Signature []byte `protobuf:"bytes,5,opt,name=signature,proto3" json:"signature,omitempty"`
	Arg       []byte `protobuf:"bytes,6,opt,name=arg,proto3" json:"arg,omitempty"`
}

func (m *CommandRequest) Reset()                    { *m = CommandRequest{} }
func (m *CommandRequest) String() string            { return proto.CompactTextString(m) }
func (*CommandRequest) ProtoMessage()               {}
func (*CommandRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CommandRequest) GetGroup() uint64 {
	if m != nil {
		return m.Group
	}
	return 0
}

func (m *CommandRequest) GetClientId() uint64 {
	if m != nil {
		return m.ClientId
	}
	return 0
}

func (m *CommandRequest) GetRequestId() uint64 {
	if m != nil {
		return m.RequestId
	}
	return 0
}

func (m *CommandRequest) GetFuncId() uint64 {
	if m != nil {
		return m.FuncId
	}
	return 0
}

func (m *CommandRequest) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *CommandRequest) GetArg() []byte {
	if m != nil {
		return m.Arg
	}
	return nil
}

type CommandResponse struct {
	Group     uint64 `protobuf:"varint,1,opt,name=group" json:"group,omitempty"`
	LeaderId  uint64 `protobuf:"varint,2,opt,name=leader_id,json=leaderId" json:"leader_id,omitempty"`
	NodeId    uint64 `protobuf:"varint,3,opt,name=node_id,json=nodeId" json:"node_id,omitempty"`
	RequestId uint64 `protobuf:"varint,4,opt,name=request_id,json=requestId" json:"request_id,omitempty"`
	Signature []byte `protobuf:"bytes,5,opt,name=signature,proto3" json:"signature,omitempty"`
	Result    []byte `protobuf:"bytes,6,opt,name=result,proto3" json:"result,omitempty"`
}

func (m *CommandResponse) Reset()                    { *m = CommandResponse{} }
func (m *CommandResponse) String() string            { return proto.CompactTextString(m) }
func (*CommandResponse) ProtoMessage()               {}
func (*CommandResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CommandResponse) GetGroup() uint64 {
	if m != nil {
		return m.Group
	}
	return 0
}

func (m *CommandResponse) GetLeaderId() uint64 {
	if m != nil {
		return m.LeaderId
	}
	return 0
}

func (m *CommandResponse) GetNodeId() uint64 {
	if m != nil {
		return m.NodeId
	}
	return 0
}

func (m *CommandResponse) GetRequestId() uint64 {
	if m != nil {
		return m.RequestId
	}
	return 0
}

func (m *CommandResponse) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *CommandResponse) GetResult() []byte {
	if m != nil {
		return m.Result
	}
	return nil
}

type LogEntry struct {
	Term    uint64          `protobuf:"varint,1,opt,name=term" json:"term,omitempty"`
	Hash    []byte          `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
	Command *CommandRequest `protobuf:"bytes,3,opt,name=command" json:"command,omitempty"`
}

func (m *LogEntry) Reset()                    { *m = LogEntry{} }
func (m *LogEntry) String() string            { return proto.CompactTextString(m) }
func (*LogEntry) ProtoMessage()               {}
func (*LogEntry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *LogEntry) GetTerm() uint64 {
	if m != nil {
		return m.Term
	}
	return 0
}

func (m *LogEntry) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *LogEntry) GetCommand() *CommandRequest {
	if m != nil {
		return m.Command
	}
	return nil
}

type RequestVoteRequest struct {
	Group       uint64 `protobuf:"varint,1,opt,name=group" json:"group,omitempty"`
	Term        uint64 `protobuf:"varint,2,opt,name=term" json:"term,omitempty"`
	LogIndex    uint64 `protobuf:"varint,3,opt,name=log_index,json=logIndex" json:"log_index,omitempty"`
	LogTerm     uint64 `protobuf:"varint,4,opt,name=log_term,json=logTerm" json:"log_term,omitempty"`
	CandidateId uint64 `protobuf:"varint,5,opt,name=candidate_id,json=candidateId" json:"candidate_id,omitempty"`
	Signature   []byte `protobuf:"bytes,6,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (m *RequestVoteRequest) Reset()                    { *m = RequestVoteRequest{} }
func (m *RequestVoteRequest) String() string            { return proto.CompactTextString(m) }
func (*RequestVoteRequest) ProtoMessage()               {}
func (*RequestVoteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *RequestVoteRequest) GetGroup() uint64 {
	if m != nil {
		return m.Group
	}
	return 0
}

func (m *RequestVoteRequest) GetTerm() uint64 {
	if m != nil {
		return m.Term
	}
	return 0
}

func (m *RequestVoteRequest) GetLogIndex() uint64 {
	if m != nil {
		return m.LogIndex
	}
	return 0
}

func (m *RequestVoteRequest) GetLogTerm() uint64 {
	if m != nil {
		return m.LogTerm
	}
	return 0
}

func (m *RequestVoteRequest) GetCandidateId() uint64 {
	if m != nil {
		return m.CandidateId
	}
	return 0
}

func (m *RequestVoteRequest) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

type RequestVoteResponse struct {
	Group       uint64 `protobuf:"varint,1,opt,name=group" json:"group,omitempty"`
	Term        uint64 `protobuf:"varint,2,opt,name=term" json:"term,omitempty"`
	NodeId      uint64 `protobuf:"varint,3,opt,name=node_id,json=nodeId" json:"node_id,omitempty"`
	CandidateId uint64 `protobuf:"varint,4,opt,name=candidate_id,json=candidateId" json:"candidate_id,omitempty"`
	Granted     bool   `protobuf:"varint,5,opt,name=granted" json:"granted,omitempty"`
	Signature   []byte `protobuf:"bytes,6,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (m *RequestVoteResponse) Reset()                    { *m = RequestVoteResponse{} }
func (m *RequestVoteResponse) String() string            { return proto.CompactTextString(m) }
func (*RequestVoteResponse) ProtoMessage()               {}
func (*RequestVoteResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *RequestVoteResponse) GetGroup() uint64 {
	if m != nil {
		return m.Group
	}
	return 0
}

func (m *RequestVoteResponse) GetTerm() uint64 {
	if m != nil {
		return m.Term
	}
	return 0
}

func (m *RequestVoteResponse) GetNodeId() uint64 {
	if m != nil {
		return m.NodeId
	}
	return 0
}

func (m *RequestVoteResponse) GetCandidateId() uint64 {
	if m != nil {
		return m.CandidateId
	}
	return 0
}

func (m *RequestVoteResponse) GetGranted() bool {
	if m != nil {
		return m.Granted
	}
	return false
}

func (m *RequestVoteResponse) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

type AppendEntriesRequest struct {
	Group        uint64                 `protobuf:"varint,1,opt,name=group" json:"group,omitempty"`
	Term         uint64                 `protobuf:"varint,2,opt,name=term" json:"term,omitempty"`
	LeaderId     uint64                 `protobuf:"varint,3,opt,name=leader_id,json=leaderId" json:"leader_id,omitempty"`
	PrevLogIndex uint64                 `protobuf:"varint,4,opt,name=prev_log_index,json=prevLogIndex" json:"prev_log_index,omitempty"`
	PrevLogTerm  uint64                 `protobuf:"varint,5,opt,name=prev_log_term,json=prevLogTerm" json:"prev_log_term,omitempty"`
	Signature    []byte                 `protobuf:"bytes,6,opt,name=signature,proto3" json:"signature,omitempty"`
	QuorumVotes  []*RequestVoteResponse `protobuf:"bytes,7,rep,name=quorum_votes,json=quorumVotes" json:"quorum_votes,omitempty"`
	Entries      []*LogEntry            `protobuf:"bytes,8,rep,name=entries" json:"entries,omitempty"`
}

func (m *AppendEntriesRequest) Reset()                    { *m = AppendEntriesRequest{} }
func (m *AppendEntriesRequest) String() string            { return proto.CompactTextString(m) }
func (*AppendEntriesRequest) ProtoMessage()               {}
func (*AppendEntriesRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *AppendEntriesRequest) GetGroup() uint64 {
	if m != nil {
		return m.Group
	}
	return 0
}

func (m *AppendEntriesRequest) GetTerm() uint64 {
	if m != nil {
		return m.Term
	}
	return 0
}

func (m *AppendEntriesRequest) GetLeaderId() uint64 {
	if m != nil {
		return m.LeaderId
	}
	return 0
}

func (m *AppendEntriesRequest) GetPrevLogIndex() uint64 {
	if m != nil {
		return m.PrevLogIndex
	}
	return 0
}

func (m *AppendEntriesRequest) GetPrevLogTerm() uint64 {
	if m != nil {
		return m.PrevLogTerm
	}
	return 0
}

func (m *AppendEntriesRequest) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *AppendEntriesRequest) GetQuorumVotes() []*RequestVoteResponse {
	if m != nil {
		return m.QuorumVotes
	}
	return nil
}

func (m *AppendEntriesRequest) GetEntries() []*LogEntry {
	if m != nil {
		return m.Entries
	}
	return nil
}

type AppendEntriesResponse struct {
	Group     uint64 `protobuf:"varint,1,opt,name=group" json:"group,omitempty"`
	Term      uint64 `protobuf:"varint,2,opt,name=term" json:"term,omitempty"`
	Index     uint64 `protobuf:"varint,3,opt,name=index" json:"index,omitempty"`
	NodeId    uint64 `protobuf:"varint,4,opt,name=node_id,json=nodeId" json:"node_id,omitempty"`
	Successed bool   `protobuf:"varint,5,opt,name=successed" json:"successed,omitempty"`
	Convinced bool   `protobuf:"varint,6,opt,name=convinced" json:"convinced,omitempty"`
	Hash      []byte `protobuf:"bytes,7,opt,name=hash,proto3" json:"hash,omitempty"`
	Signature []byte `protobuf:"bytes,8,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (m *AppendEntriesResponse) Reset()                    { *m = AppendEntriesResponse{} }
func (m *AppendEntriesResponse) String() string            { return proto.CompactTextString(m) }
func (*AppendEntriesResponse) ProtoMessage()               {}
func (*AppendEntriesResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *AppendEntriesResponse) GetGroup() uint64 {
	if m != nil {
		return m.Group
	}
	return 0
}

func (m *AppendEntriesResponse) GetTerm() uint64 {
	if m != nil {
		return m.Term
	}
	return 0
}

func (m *AppendEntriesResponse) GetIndex() uint64 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *AppendEntriesResponse) GetNodeId() uint64 {
	if m != nil {
		return m.NodeId
	}
	return 0
}

func (m *AppendEntriesResponse) GetSuccessed() bool {
	if m != nil {
		return m.Successed
	}
	return false
}

func (m *AppendEntriesResponse) GetConvinced() bool {
	if m != nil {
		return m.Convinced
	}
	return false
}

func (m *AppendEntriesResponse) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *AppendEntriesResponse) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

type Peer struct {
	Id   uint64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Host uint64 `protobuf:"varint,2,opt,name=host" json:"host,omitempty"`
}

func (m *Peer) Reset()                    { *m = Peer{} }
func (m *Peer) String() string            { return proto.CompactTextString(m) }
func (*Peer) ProtoMessage()               {}
func (*Peer) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *Peer) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Peer) GetHost() uint64 {
	if m != nil {
		return m.Host
	}
	return 0
}

type Node struct {
	Id         uint64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	LastSeen   uint64 `protobuf:"varint,2,opt,name=last_seen,json=lastSeen" json:"last_seen,omitempty"`
	Weight     uint32 `protobuf:"varint,3,opt,name=weight" json:"weight,omitempty"`
	Online     bool   `protobuf:"varint,4,opt,name=online" json:"online,omitempty"`
	PublicKey  string `protobuf:"bytes,5,opt,name=public_key,json=publicKey" json:"public_key,omitempty"`
	ServerAddr string `protobuf:"bytes,6,opt,name=server_addr,json=serverAddr" json:"server_addr,omitempty"`
}

func (m *Node) Reset()                    { *m = Node{} }
func (m *Node) String() string            { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()               {}
func (*Node) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *Node) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Node) GetLastSeen() uint64 {
	if m != nil {
		return m.LastSeen
	}
	return 0
}

func (m *Node) GetWeight() uint32 {
	if m != nil {
		return m.Weight
	}
	return 0
}

func (m *Node) GetOnline() bool {
	if m != nil {
		return m.Online
	}
	return false
}

func (m *Node) GetPublicKey() string {
	if m != nil {
		return m.PublicKey
	}
	return ""
}

func (m *Node) GetServerAddr() string {
	if m != nil {
		return m.ServerAddr
	}
	return ""
}

type RaftGroup struct {
	Replications uint32 `protobuf:"varint,1,opt,name=replications" json:"replications,omitempty"`
	Id           uint64 `protobuf:"varint,2,opt,name=id" json:"id,omitempty"`
	LeaderPeer   uint64 `protobuf:"varint,3,opt,name=leader_peer,json=leaderPeer" json:"leader_peer,omitempty"`
}

func (m *RaftGroup) Reset()                    { *m = RaftGroup{} }
func (m *RaftGroup) String() string            { return proto.CompactTextString(m) }
func (*RaftGroup) ProtoMessage()               {}
func (*RaftGroup) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *RaftGroup) GetReplications() uint32 {
	if m != nil {
		return m.Replications
	}
	return 0
}

func (m *RaftGroup) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *RaftGroup) GetLeaderPeer() uint64 {
	if m != nil {
		return m.LeaderPeer
	}
	return 0
}

func init() {
	proto.RegisterType((*CommandRequest)(nil), "bftraft.CommandRequest")
	proto.RegisterType((*CommandResponse)(nil), "bftraft.CommandResponse")
	proto.RegisterType((*LogEntry)(nil), "bftraft.LogEntry")
	proto.RegisterType((*RequestVoteRequest)(nil), "bftraft.RequestVoteRequest")
	proto.RegisterType((*RequestVoteResponse)(nil), "bftraft.RequestVoteResponse")
	proto.RegisterType((*AppendEntriesRequest)(nil), "bftraft.AppendEntriesRequest")
	proto.RegisterType((*AppendEntriesResponse)(nil), "bftraft.AppendEntriesResponse")
	proto.RegisterType((*Peer)(nil), "bftraft.Peer")
	proto.RegisterType((*Node)(nil), "bftraft.Node")
	proto.RegisterType((*RaftGroup)(nil), "bftraft.RaftGroup")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for BFTRaft service

type BFTRaftClient interface {
	ExecCommand(ctx context.Context, in *CommandRequest, opts ...grpc.CallOption) (*CommandResponse, error)
	RequestVote(ctx context.Context, in *RequestVoteRequest, opts ...grpc.CallOption) (*RequestVoteResponse, error)
	AppendEntries(ctx context.Context, in *AppendEntriesRequest, opts ...grpc.CallOption) (*AppendEntriesResponse, error)
}

type bFTRaftClient struct {
	cc *grpc.ClientConn
}

func NewBFTRaftClient(cc *grpc.ClientConn) BFTRaftClient {
	return &bFTRaftClient{cc}
}

func (c *bFTRaftClient) ExecCommand(ctx context.Context, in *CommandRequest, opts ...grpc.CallOption) (*CommandResponse, error) {
	out := new(CommandResponse)
	err := grpc.Invoke(ctx, "/bftraft.BFTRaft/ExecCommand", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bFTRaftClient) RequestVote(ctx context.Context, in *RequestVoteRequest, opts ...grpc.CallOption) (*RequestVoteResponse, error) {
	out := new(RequestVoteResponse)
	err := grpc.Invoke(ctx, "/bftraft.BFTRaft/RequestVote", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bFTRaftClient) AppendEntries(ctx context.Context, in *AppendEntriesRequest, opts ...grpc.CallOption) (*AppendEntriesResponse, error) {
	out := new(AppendEntriesResponse)
	err := grpc.Invoke(ctx, "/bftraft.BFTRaft/AppendEntries", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for BFTRaft service

type BFTRaftServer interface {
	ExecCommand(context.Context, *CommandRequest) (*CommandResponse, error)
	RequestVote(context.Context, *RequestVoteRequest) (*RequestVoteResponse, error)
	AppendEntries(context.Context, *AppendEntriesRequest) (*AppendEntriesResponse, error)
}

func RegisterBFTRaftServer(s *grpc.Server, srv BFTRaftServer) {
	s.RegisterService(&_BFTRaft_serviceDesc, srv)
}

func _BFTRaft_ExecCommand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BFTRaftServer).ExecCommand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bftraft.BFTRaft/ExecCommand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BFTRaftServer).ExecCommand(ctx, req.(*CommandRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BFTRaft_RequestVote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestVoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BFTRaftServer).RequestVote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bftraft.BFTRaft/RequestVote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BFTRaftServer).RequestVote(ctx, req.(*RequestVoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BFTRaft_AppendEntries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppendEntriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BFTRaftServer).AppendEntries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bftraft.BFTRaft/AppendEntries",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BFTRaftServer).AppendEntries(ctx, req.(*AppendEntriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BFTRaft_serviceDesc = grpc.ServiceDesc{
	ServiceName: "bftraft.BFTRaft",
	HandlerType: (*BFTRaftServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ExecCommand",
			Handler:    _BFTRaft_ExecCommand_Handler,
		},
		{
			MethodName: "RequestVote",
			Handler:    _BFTRaft_RequestVote_Handler,
		},
		{
			MethodName: "AppendEntries",
			Handler:    _BFTRaft_AppendEntries_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/bftraft.proto",
}

func init() { proto.RegisterFile("proto/bftraft.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 763 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0xdd, 0x8e, 0xdb, 0x44,
	0x14, 0xae, 0x13, 0x27, 0x76, 0x8e, 0x93, 0x05, 0xa6, 0x0b, 0x35, 0xfd, 0xa1, 0x8b, 0xc5, 0xc5,
	0x0a, 0xa4, 0x22, 0x96, 0x07, 0x40, 0x2d, 0x2a, 0x28, 0x50, 0xa1, 0xca, 0x54, 0xdc, 0x1a, 0xaf,
	0xe7, 0xc4, 0x6b, 0xe1, 0xcc, 0xb8, 0x33, 0xe3, 0xd0, 0x7d, 0x1e, 0x2e, 0xb9, 0x83, 0x0b, 0xde,
	0x86, 0x17, 0xe0, 0x25, 0xd0, 0xfc, 0xd8, 0xf9, 0x0f, 0xa2, 0x77, 0x73, 0xbe, 0x33, 0x33, 0xf9,
	0xce, 0xf7, 0x7d, 0x9e, 0xc0, 0xdd, 0x46, 0x70, 0xc5, 0x3f, 0xbf, 0x5e, 0x28, 0x91, 0x2f, 0xd4,
	0x13, 0x53, 0x91, 0xc0, 0x95, 0xc9, 0xef, 0x1e, 0x9c, 0x7d, 0xcd, 0x97, 0xcb, 0x9c, 0xd1, 0x14,
	0x5f, 0xb7, 0x28, 0x15, 0x39, 0x87, 0x51, 0x29, 0x78, 0xdb, 0xc4, 0xde, 0x85, 0x77, 0xe9, 0xa7,
	0xb6, 0x20, 0x0f, 0x60, 0x52, 0xd4, 0x15, 0x32, 0x95, 0x55, 0x34, 0x1e, 0x98, 0x4e, 0x68, 0x81,
	0x39, 0x25, 0x8f, 0x00, 0x84, 0x3d, 0xad, 0xbb, 0x43, 0xd3, 0x9d, 0x38, 0x64, 0x4e, 0xc9, 0x3d,
	0x08, 0x16, 0x2d, 0x2b, 0x74, 0xcf, 0x37, 0xbd, 0xb1, 0x2e, 0xe7, 0x94, 0x3c, 0x84, 0x89, 0xac,
	0x4a, 0x96, 0xab, 0x56, 0x60, 0x3c, 0xba, 0xf0, 0x2e, 0xa7, 0xe9, 0x1a, 0x20, 0xef, 0xc2, 0x30,
	0x17, 0x65, 0x3c, 0x36, 0xb8, 0x5e, 0x26, 0x7f, 0x78, 0xf0, 0x4e, 0xcf, 0x56, 0x36, 0x9c, 0x49,
	0x3c, 0x4e, 0xb7, 0xc6, 0x9c, 0xa2, 0xd8, 0xa0, 0x6b, 0x01, 0xcb, 0x87, 0x71, 0x8a, 0x6b, 0xae,
	0x63, 0x5d, 0xee, 0xcd, 0xe1, 0xef, 0xce, 0x71, 0x9a, 0xee, 0x07, 0x30, 0x16, 0x28, 0xdb, 0x5a,
	0x39, 0xc6, 0xae, 0x4a, 0x10, 0xc2, 0x17, 0xbc, 0x7c, 0xce, 0x94, 0xb8, 0x25, 0x04, 0x7c, 0x85,
	0x62, 0xe9, 0xb8, 0x9a, 0xb5, 0xc6, 0x6e, 0x72, 0x79, 0x63, 0x58, 0x4e, 0x53, 0xb3, 0x26, 0x5f,
	0x40, 0x50, 0xd8, 0x39, 0x0d, 0xc3, 0xe8, 0xea, 0xde, 0x93, 0xce, 0xc0, 0x6d, 0xb7, 0xd2, 0x6e,
	0x5f, 0xf2, 0x97, 0x07, 0xc4, 0x81, 0x3f, 0x71, 0x85, 0xa7, 0xdd, 0xec, 0x78, 0x0c, 0x36, 0x78,
	0x68, 0xc9, 0x78, 0x99, 0x55, 0x8c, 0xe2, 0x1b, 0xa7, 0x4b, 0x58, 0xf3, 0x72, 0xae, 0x6b, 0xf2,
	0x21, 0xe8, 0x75, 0x66, 0x0e, 0x59, 0x5d, 0x82, 0x9a, 0x97, 0xaf, 0xf4, 0xb9, 0x8f, 0x61, 0x5a,
	0xe4, 0x8c, 0x56, 0x34, 0x57, 0x46, 0xd2, 0x91, 0x69, 0x47, 0x3d, 0xb6, 0x2b, 0xdc, 0x78, 0x47,
	0xb8, 0xe4, 0x4f, 0x0f, 0xee, 0x6e, 0x31, 0x3f, 0xe9, 0xec, 0x21, 0xea, 0x47, 0x0d, 0xdd, 0xe5,
	0xe6, 0xef, 0x73, 0x8b, 0x21, 0x28, 0x45, 0xce, 0x14, 0x5a, 0xe6, 0x61, 0xda, 0x95, 0xff, 0xc5,
	0x7a, 0x00, 0xe7, 0x4f, 0x9b, 0x06, 0x19, 0xd5, 0xd6, 0x56, 0x28, 0xdf, 0x4e, 0xf1, 0x3e, 0xa4,
	0xc3, 0x9d, 0x90, 0x7e, 0x02, 0x67, 0x8d, 0xc0, 0x55, 0xb6, 0xf6, 0xc4, 0x92, 0x9f, 0x6a, 0xf4,
	0x45, 0xe7, 0x4b, 0x02, 0xb3, 0x7e, 0x97, 0xb9, 0xdf, 0xa9, 0xef, 0x36, 0x19, 0x83, 0x4e, 0xce,
	0x41, 0xbe, 0x82, 0xe9, 0xeb, 0x96, 0x8b, 0x76, 0x99, 0xad, 0xb8, 0x42, 0x19, 0x07, 0x17, 0xc3,
	0xcb, 0xe8, 0xea, 0x61, 0x9f, 0xb7, 0x03, 0xce, 0xa4, 0x91, 0x3d, 0xa1, 0x31, 0x49, 0x3e, 0x83,
	0x00, 0xad, 0x02, 0x71, 0x68, 0xce, 0xbe, 0xd7, 0x9f, 0xed, 0x72, 0x9f, 0x76, 0x3b, 0x92, 0xbf,
	0x3d, 0x78, 0x7f, 0x47, 0xb5, 0xff, 0xed, 0xf6, 0x39, 0x8c, 0x36, 0x43, 0x6a, 0x8b, 0xcd, 0x0c,
	0xf8, 0x5b, 0x19, 0xd0, 0xe3, 0xb7, 0x45, 0x81, 0x52, 0xf6, 0x16, 0xaf, 0x01, 0xdd, 0x2d, 0x38,
	0x5b, 0x55, 0xac, 0x40, 0x6a, 0xc4, 0x09, 0xd3, 0x35, 0xd0, 0x7f, 0x9b, 0xc1, 0xc6, 0xb7, 0xb9,
	0x25, 0x67, 0xb8, 0x1b, 0x8b, 0x4f, 0xc1, 0x7f, 0x89, 0x28, 0xc8, 0x19, 0x0c, 0x2a, 0xea, 0x66,
	0x19, 0x54, 0xf6, 0x26, 0x2e, 0x55, 0x37, 0x88, 0x5e, 0x27, 0xbf, 0x79, 0xe0, 0xff, 0xc0, 0x29,
	0xee, 0x6d, 0xd6, 0xc1, 0xc8, 0xa5, 0xca, 0x24, 0x22, 0xeb, 0x5f, 0xaf, 0x5c, 0xaa, 0x1f, 0x11,
	0x99, 0x7e, 0x67, 0x7e, 0xc5, 0xaa, 0xbc, 0x51, 0x66, 0xfe, 0x59, 0xea, 0x2a, 0x8d, 0x73, 0x56,
	0x57, 0x0c, 0xcd, 0xfc, 0x61, 0xea, 0x2a, 0xfd, 0xa8, 0x35, 0xed, 0x75, 0x5d, 0x15, 0xd9, 0x2f,
	0x78, 0x6b, 0x04, 0x98, 0xa4, 0x13, 0x8b, 0x7c, 0x8f, 0xb7, 0xe4, 0x31, 0x44, 0x12, 0xc5, 0x0a,
	0x45, 0x96, 0x53, 0x2a, 0x8c, 0x04, 0x93, 0x14, 0x2c, 0xf4, 0x94, 0x52, 0x91, 0xfc, 0x0c, 0x93,
	0x34, 0x5f, 0xa8, 0x6f, 0x8d, 0x1f, 0x09, 0x4c, 0x05, 0x36, 0x75, 0x55, 0xe4, 0xaa, 0xe2, 0x4c,
	0x1a, 0xce, 0xb3, 0x74, 0x0b, 0x73, 0xd3, 0x0c, 0xfa, 0x69, 0x1e, 0x43, 0xe4, 0x62, 0xde, 0x20,
	0x0a, 0xe7, 0x1a, 0x58, 0x48, 0x6b, 0x75, 0xf5, 0x8f, 0x07, 0xc1, 0xb3, 0x6f, 0x5e, 0xe9, 0x5f,
	0x21, 0xcf, 0x20, 0x7a, 0xfe, 0x06, 0x0b, 0xf7, 0xca, 0x91, 0x63, 0xef, 0xde, 0xfd, 0x78, 0xbf,
	0x61, 0x83, 0x94, 0xdc, 0x21, 0xdf, 0x41, 0xb4, 0x91, 0x5a, 0xf2, 0xe0, 0x70, 0x96, 0xed, 0x3d,
	0x27, 0x83, 0x9e, 0xdc, 0x21, 0x2f, 0x61, 0xb6, 0x95, 0x57, 0xf2, 0xa8, 0x3f, 0x70, 0xe8, 0xeb,
	0xbf, 0xff, 0xd1, 0xb1, 0x76, 0x77, 0xe3, 0xf5, 0xd8, 0xfc, 0x05, 0x7f, 0xf9, 0x6f, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xdd, 0x53, 0xa0, 0xb7, 0x99, 0x07, 0x00, 0x00,
}
