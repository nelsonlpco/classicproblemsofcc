package fib

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var fibTestTable = []struct {
	v uint64
	r uint64
}{
	{v: 1, r: 1},
	{v: 2, r: 1},
	{v: 3, r: 2},
	{v: 6, r: 8},
	{v: 40, r: 102334155},
}

func Test_FibonacciSequential(t *testing.T) {
	for _, test := range fibTestTable {
		t.Run(fmt.Sprintf("FibonacciSequential: %v = %v", test.v, test.r), func(t *testing.T) {
			result := FibonacciSequential(test.v)
			fmt.Println(result)
			assert.Equal(t, test.r, result)
		})
	}
}

func Test_FibonacciRec(t *testing.T) {
	for _, test := range fibTestTable {
		t.Run(fmt.Sprintf("FibonacciRec: %v = %v", test.v, test.r), func(t *testing.T) {
			result := FibonacciRec(test.v)
			fmt.Println(result)
			assert.Equal(t, test.r, result)
		})
	}
}

func Test_FibonacciRecCached(t *testing.T) {
	for _, test := range fibTestTable {
		t.Run(fmt.Sprintf("FibonacciRecCached: %v = %v", test.v, test.r), func(t *testing.T) {
			result := FibonacciRecCached(test.v)
			fmt.Println(result)
			assert.Equal(t, test.r, result)
		})
	}
}

func Test_MemoizeFib(t *testing.T) {
	mFib := MemoizeFib(FibonacciRecCached)
	for _, test := range fibTestTable {
		t.Run(fmt.Sprintf("FibonacciRecCached: %v = %v", test.v, test.r), func(t *testing.T) {
			result := mFib(test.v)
			fmt.Println(result)
			assert.Equal(t, test.r, result)
		})
	}
}

func Test_GenerateFib(t *testing.T) {
	expectedValues := []uint64{1, 1, 2, 3, 5, 8}
	i := 0
	for v := range GenerateSequentialFibonacci(6) {
		assert.Equal(t, expectedValues[i], v)
		i++
	}
}

var table = []struct {
	input uint64
}{
	{input: 30},
	{input: 80},
	{input: 20},
	{input: 80},
	{input: 20},
	{input: 30},
}

func Benchmark_CachedFib(b *testing.B) {
	mem = make(map[uint64]uint64)
	for _, v := range table {
		for i := 0; i < b.N; i++ {
			FibonacciRecCached(v.input)
		}
	}
}

func Benchmark_MemoizedFib(b *testing.B) {
	mem = make(map[uint64]uint64)
	mFib := MemoizeFib(FibonacciRecCached)

	for _, v := range table {
		for i := 0; i < b.N; i++ {
			mFib(v.input)
		}
	}
}

func Benchmark_SequentialFibonacci(b *testing.B) {
	for _, v := range table {
		for i := 0; i < b.N; i++ {
			FibonacciSequential(v.input)
		}
	}
}

func Benchmark_MemoizedSequentialFibonacci(b *testing.B) {
	mem = make(map[uint64]uint64)
	mFib := MemoizeFib(FibonacciSequential)
	for _, v := range table {
		for i := 0; i < b.N; i++ {
			mFib(v.input)
		}
	}
}
