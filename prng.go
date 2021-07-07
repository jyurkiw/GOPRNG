package goprng

/** Seed a 320-bit xorshift psuedo random number
generator function that creates int64 values
and return it. */
func seedXOrShift64(a int64, b int64, c int64, d int64, e int64) func() int64 {
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

/** An int64 prng interface containing definitions for
next, next with a ceiling, and next between a ceiling
and floor functions.
*/
type Prng interface {
	Next() int64
	NextMax(max int64) int64
	NextBetween(min int64, max int64) int64
}

/** A basic xorshift struct.
Contains a generator function.
*/
type xorshift struct {
	generator func() int64
}

/** Return a psuedo-random int64.
 */
func (g xorshift) Next() int64 {
	return g.generator()
}

/** Return a psuedo-random int64 with a value
of 0 inclusive < value < max exclusive.
*/
func (g xorshift) NextMax(max int64) int64 {
	return g.generator() % max
}

/** Return a psuedo-random int64 with a value
of min inclusive < value < max exclusive.
*/
func (g xorshift) NextBetween(min int64, max int64) int64 {
	return (g.generator() % (max - min)) + min
}

/** Seed and return an xorshift prng.
 */
func New(a int64, b int64, c int64, d int64, e int64) Prng {
	return xorshift{generator: seedXOrShift64(a, b, c, d, e)}
}
