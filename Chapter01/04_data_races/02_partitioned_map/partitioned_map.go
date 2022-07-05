package _2_partitioned_map

import (
	"hash/fnv"
	"sync"
)

func NewPartitionedMap(partitions int) *PartitionedMap {
	out := &PartitionedMap{
		partitions: make([]*partition, partitions),
	}

	for x := 0; x < partitions; x++ {
		out.partitions[x] = &partition{
			data:  map[string]interface{}{},
			mutex: sync.RWMutex{},
		}
	}

	return out
}

type PartitionedMap struct {
	partitions []*partition
}

type partition struct {
	data  map[string]interface{}
	mutex sync.RWMutex
}

func (p *PartitionedMap) Get(key string) (interface{}, bool) {
	partition := p.getPartition(key)

	partition.mutex.RLock()
	defer partition.mutex.RUnlock()

	value, exists := partition.data[key]
	return value, exists
}

func (p *PartitionedMap) Set(key string, value interface{}) {
	partition := p.getPartition(key)

	partition.mutex.Lock()
	defer partition.mutex.Unlock()

	partition.data[key] = value
}

func (p *PartitionedMap) getPartition(key string) *partition {
	// calculate partition index from key using the FNV-1
	// from the standard library
	hashGenerator := fnv.New32()
	_, _ = hashGenerator.Write([]byte(key))

	partitionIndex := int(hashGenerator.Sum32()) % len(p.partitions)

	return p.partitions[partitionIndex]
}
