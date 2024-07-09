package gevm

import (
	"github.com/ethereum/go-ethereum/common"
)

type Storage struct {
	data  map[int]common.Hash
	cache map[int]bool // slot keys cache for warm storage access
}

func (s *Storage) Load(key int) (value common.Hash, isWarm bool) {
	isWarm = s.cache[key]
	if !isWarm {
		s.cache[key] = true
	}
	value, ok := s.data[key]
	if !ok {
		return common.Hash{}, false
	}
	return value, ok
}

func (s *Storage) Store(key int, value common.Hash) (isWarm bool) {
	isWarm = s.cache[key]
	if !isWarm {
		s.cache[key] = true
	}
	s.data[key] = value
	return isWarm
}

// Get does the same thing as Load, except that it doesn't mark the storage slot as 'warm'.
// It is used in the 'calcSstoreGasCost' function in 'common.go'
func (s *Storage) Get(slot int) (value common.Hash, isWarm bool) {
	return s.data[slot], s.cache[slot]
}

func NewStorage() *Storage {
	return &Storage{
		cache: make(map[int]bool),
		data:  make(map[int]common.Hash),
	}
}
