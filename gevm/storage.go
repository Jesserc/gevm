package gevm

import (
	"slices"

	"github.com/ethereum/go-ethereum/common"
)

type Storage struct {
	data  map[int]common.Hash
	cache []int // slot keys cache for warm storage access
}

func (s *Storage) Load(key int) (isWarm bool, _ common.Hash) {
	isWarm = slices.Contains(s.cache, key)
	if !isWarm {
		s.cache = append(s.cache, key)
	}
	if _, ok := s.data[key]; !ok {
		return false, common.HexToHash("0x00")
	}
	return isWarm, s.data[key]
}

func (s *Storage) Store(key int, value common.Hash) (isWarm bool) {
	prev := s.data[key]
	s.data[key] = value
	isWarm = prev == value
	return isWarm
}
func NewStorage() *Storage {
	return &Storage{
		data:  make(map[int]common.Hash),
		cache: make([]int, 0),
	}
}
