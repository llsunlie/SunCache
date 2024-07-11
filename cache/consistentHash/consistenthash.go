package consistentHash

import (
	"hash"
	"sort"
	"strconv"

	"github.com/spaolacci/murmur3"
)

const DefaultMultiple = 50

type ConsistentHash struct {
	hash     hash.Hash32
	multiple int
	nodes    []uint32
	hashMap  map[uint32]string
}

func (c *ConsistentHash) Hash32(b []byte) (hash uint32) {
	c.hash.Write(b)
	hash = c.hash.Sum32()
	c.hash.Reset()
	return
}

func NewConsistentHash(multiple int, hash hash.Hash32) (consistentHash *ConsistentHash) {
	consistentHash = &ConsistentHash{
		hash:     hash,
		multiple: multiple,
		hashMap:  make(map[uint32]string),
	}
	if consistentHash.hash == nil {
		consistentHash.hash = murmur3.New32()
	}
	return
}

func (c *ConsistentHash) Add(nodes ...string) {
	for _, node := range nodes {
		for i := 0; i < c.multiple; i++ {
			hash := c.Hash32([]byte(strconv.Itoa(i) + node))
			c.nodes = append(c.nodes, hash)
			c.hashMap[hash] = node
		}
	}
	sort.Slice(c.nodes, func(i, j int) bool {
		return c.nodes[i] < c.nodes[j]
	})
}

func (c *ConsistentHash) Get(key string) (node string) {
	if len(c.nodes) == 0 {
		return ""
	}
	hash := c.Hash32([]byte(key))
	idx := sort.Search(len(c.nodes), func(i int) bool {
		return c.nodes[i] >= hash
	})
	return c.hashMap[c.nodes[idx%len(c.nodes)]]
}
