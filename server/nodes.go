package server

import (
	"github.com/dgraph-io/badger"
	"github.com/golang/protobuf/proto"
	pb "github.com/PomeloCloud/BFTRaft4go/proto"
	jump "github.com/renstrom/go-jump-consistent-hash"
	"hash/fnv"

)

type NodeIterator struct {
	prefix []byte
	data *badger.Iterator
}

func (liter *NodeIterator) Next() *pb.Node  {
	liter.data.Next()
	if liter.data.ValidForPrefix(liter.prefix) {
		entry := pb.Node{}
		itemData := ItemValue(liter.data.Item())
		proto.Unmarshal(itemData, &entry)
		return &entry
	} else {
		return nil
	}
}

func (s *BFTRaftServer) NodesIterator () NodeIterator {
	keyPrefix := ComposeKeyPrefix(NODE_LIST_GROUP, NODES_LIST)
	iter := s.DB.NewIterator(badger.IteratorOptions{})
	iter.Seek(append(keyPrefix, U64Bytes(0)...))
	return NodeIterator{
		prefix: keyPrefix,
		data: iter,
	}
}

func (s *BFTRaftServer) LoadOnlineNodes()  {
	iter := s.NodesIterator()
	nodes := map[uint64]*pb.Node{}
	for true {
		if node := iter.Next(); node != nil {
			if node.Online {
				nodes[node.Id] = node
			}
		} else {
			break
		}
	}
	s.Nodes = nodes
}

func (s *BFTRaftServer) LocateNodeIndex(keyword []byte) int32 {
	h := fnv.New64a()
	h.Write(keyword)
	sum := h.Sum64()
	return jump.Hash(sum, int32(len(s.Nodes)))
}

