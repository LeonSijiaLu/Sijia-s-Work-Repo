// 一致性 Hash 实现

package main

import (
	"crypto/md5"
	"fmt"
	"math"
	"sort"
)

type HashKey uint32
type HashKeyOrder []HashKey

func (h HashKeyOrder) Len() int {return len(h)}
func (h HashKeyOrder) Swap(i int, j int) {h[i], h[j] = h[j], h[i]}
func (h HashKeyOrder) Less(i int, j int) bool {return h[i] < h[j]}

type HashRing struct{
	ring map[HashKey]string
	sortedKeys []HashKey
	nodes []string
	weights map[string]int
}

func (hashRing *HashRing) generateCycle(){
	
}

func New(nodes []string) *HashRing{
	hashRing := &HashRing{
		ring: make(map[HashKey] string),
		sortedKeys: make([]HashKey, 0),
		nodes: nodes,
		weights: make(map[string] int),
	}
	hashRing.generateCycle()
	return hashRing
}