package gevm

import (
	"slices"

	"github.com/ethereum/go-ethereum/common"
)

type Storage struct {
	data  map[int]common.Hash
	cache []int // slot keys cache for warm storage access
}

func (s *Storage) Load(key int) (bool, common.Hash) {
	warmAccess := slices.Contains(s.cache, key)
	if !warmAccess {
		s.cache = append(s.cache, key)
	}
	if _, ok := s.data[key]; !ok {
		return false, common.HexToHash("0x00")
	}
	return warmAccess, s.data[key]
}

func (s *Storage) Store(key int, value common.Hash) {
	s.data[key] = value
}

func NewStorage() *Storage {
	return &Storage{
		data:  make(map[int]common.Hash),
		cache: make([]int, 0),
	}
}
