package xorshift64

func SeedXOrShift64(a int64, b int64, c int64, d int64, e int64) func() int64 {
	return func() int64 {
		x := e
		s := a

		e = d
		d = c
		c = b
		b = c

		x ^= x << 23
		x ^= x >> 17
		a = x ^ s ^ (s >> 19)
		return a
	}
}

type prng interface {
	next() int64
	nextMax(max int64) int64
	nextBetween(min int64, max int64) int64
}

type xorshift struct {
	generator func() int64
}

func (g xorshift) next() int64 {
	return g.generator()
}

func (g xorshift) nextMax(max int64) int64 {
	return g.generator() % max
}

func (g xorshift) nextBetween(min int64, max int64) int64 {
	return (g.generator() % (max - min)) + min
}

func new(a int64, b int64, c int64, d int64, e int64) prng {
	return xorshift{generator: SeedXOrShift64(a, b, c, d, e)}
}
