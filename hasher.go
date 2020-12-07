package main

import(
	"hash/fnv"
	"strconv"
)

type Hasher struct { }

type HashGenerator interface {
	GenerateHashFromURL(url string) string
}

func (h *Hasher) GenerateHashFromURL(url string) string {
	return strconv.FormatUint(uint64(h.HashString(url)), 10)
}

func (h *Hasher) HashString(s string) uint32 {
	hash := fnv.New32a()
	hash.Write([]byte(s))
	return hash.Sum32()
}