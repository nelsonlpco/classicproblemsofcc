package fib

var mem = make(map[uint64]uint64)

func FibonacciSequential(n uint64) uint64 {
	if n < 2 {
		return n
	}

	n2, n1 := uint64(0), uint64(1)
	for i := uint64(0); i < n; i++ {
		n2, n1 = n1, n1+n2
	}

	return n2
}

func FibonacciRec(n uint64) uint64 {
	if n < 2 {
		return n
	}

	return FibonacciRec(n-2) + FibonacciRec(n-1)
}

func FibonacciRecCached(n uint64) uint64 {
	if n < 2 {
		return n
	}

	if _, hasValue := mem[n]; hasValue {
		return mem[n]
	}

	mem[n] = FibonacciRecCached(n-2) + FibonacciRecCached(n-1)

	return mem[n]
}

func MemoizeFib(fib func(uint64) uint64) func(uint64) uint64 {
	cache := make(map[uint64]uint64)
	return func(n uint64) uint64 {
		if value, hasValue := cache[n]; hasValue {
			return value
		}
		cache[n] = fib(n)

		return cache[n]
	}
}

func GenerateSequentialFibonacci(n uint64) chan uint64 {
	ch := make(chan uint64)
	// sem := semaphore.NewWeighted(int64(n))

	go func() {
		if n < 2 {
			ch <- n
		}
		n1, n2 := uint64(1), uint64(0)
		for i := uint64(0); i < n; i++ {
			n2, n1 = n1, n1+n2
			ch <- n2
		}
		close(ch)
	}()

	return ch
}
