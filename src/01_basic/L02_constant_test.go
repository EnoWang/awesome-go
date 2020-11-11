package basic

import (
	"math"
	"testing"
)

// 0. const declared
/*
a. When no type specify for const, they can be any type
b. can be out of function (in package scope)
c. No need to make const names in all capital letters
   - special meaning with capital letter in GoLang
*/
func TestConst(t *testing.T) {
	const filename = "abc.txt"
	const a, b = 3, 4
	var c int
	// note: a and b are used as float
	c = int(math.Sqrt(a*a + b*b))
	t.Log(filename)
	t.Log("triangle:", c)
}

// 1. Enum
/*
a. No Enum key word, create enum by creating a bunch of constants
b. key word iota: this const array is auto-incremental
    i. can use _ to skip a value
	ii. can compose complicated expression
*/
func TestEnums(t *testing.T) {
	const (
		cpp    = 0
		java   = 1
		python = 2
		golang = 3
	)
	const (
		zero = 0
		one  = 1
		_
		three = 3
	)
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	t.Log(cpp, java, python, golang)
	t.Log(zero, one, three)
	t.Log(b, kb, mb, gb, tb, pb)
}
