package filter

import (
	"hash"
	"hash/fnv"
	"log"

	"github.com/steakknife/bloomfilter"
)

const (
	BLOOM     = "bloom"
	BLOOMSIZE = 1000000
)

type Filter interface {
	Contains(string) bool
}

type Bloom struct {
	maxElements uint64
	bf          *bloomfilter.Filter
	impactN     int
	m           map[string]int
}

func (b *Bloom) Contains(s string) bool {
	h := shash(s)
	if v, ok := b.m[s]; ok {
		b.m[s] = v + 1
	} else {
		b.m[s] = 1
	}
	if b.bf.Contains(h) {

		b.impactN++
		return true
	}
	b.impactN--
	b.bf.Add(h)
	return false
}

func shash(s string) hash.Hash64 {
	var h64 hash.Hash64 = fnv.New64()
	h64.Write([]byte(s))
	return h64
}

func New(n uint64) *Bloom {
	const probCollide = 0.0000001
	bf, err := bloomfilter.NewOptimal(n, probCollide)
	if err != nil {
		log.Panic(err.Error())
	}
	m := make(map[string]int)

	return &Bloom{
		maxElements: n,
		bf:          bf,
		m:           m,
	}
}
