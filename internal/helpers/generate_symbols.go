package helpers

import (
	"github.com/stretchr/testify/mock"
	"math/rand"
)

type IRandomSymbols interface {
	Generate() []int32
}

type RandomSymbols struct {
	Symbols []int32
	Length  int
}

func NewRandomSymbols(length int) *RandomSymbols {
	Symbols := make([]int32, length)

	return &RandomSymbols{
		Length:  length,
		Symbols: Symbols,
	}
}

func (rs *RandomSymbols) Generate() []int32 {
	for i := 0; i < 3; i++ {
		rs.Symbols[i] = rand.Int31n(9) + 1
	}

	return rs.Symbols
}

type RandomSymbolsMock struct {
	mock.Mock
}

func NewRandomSymbolsMock() *RandomSymbolsMock {
	return &RandomSymbolsMock{}
}

func (rs *RandomSymbolsMock) Generate() []int32 {
	args := rs.Called()
	return args.Get(0).([]int32)
}
