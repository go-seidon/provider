package random

type Randomizer interface {
	// @note: return n amount of random string
	// e.g: n = 5, result = sl1q0
	String(n int) (string, error)
}
