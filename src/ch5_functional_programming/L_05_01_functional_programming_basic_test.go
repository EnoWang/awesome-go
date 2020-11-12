package ch5_functional_programming

import "testing"

// 0. non-strict functional programming
// Note: closure can tough variables in environment
func adder() func(int) int {
	sum := 0
	return func(v int) int {
		sum += v
		return sum
	}
}

func TestNonStrictClosure(t *testing.T) {
	a := adder()
	for i := 0; i < 10; i++ {
		t.Log(a(i))
	}
}

// 1. strict functional programming (without local variables)
type iAdder func(int) (int, iAdder)

func adder2(base int) iAdder {
	return func(v int) (int, iAdder) {
		return base + v, adder2(base + v)
	}
}

func TestStrictClosure(t *testing.T) {
	a := adder2(0)
	for i := 0; i < 10; i++ {
		var sum int
		sum, a = a(i)
		t.Log(sum)
	}
}
