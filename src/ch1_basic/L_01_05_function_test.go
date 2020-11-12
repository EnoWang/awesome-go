package basic

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
	"testing"
)

// 0. basic function
func eval(a, b int, op string) int {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	}
	return -1
}
func TestFunc(t *testing.T) {
	t.Log(eval(1, 3, "+"))
}

// 1. return more than one value
// Note: can assign name to return values
func div(a, b int) (q, r int) {
	return a / b, a % b
}
func TestDiv(t *testing.T) {
	// No need to be q, r
	q, r := div(10, 3)
	t.Log(q, r)
}

// 2. functional programming
func apply(op func(float64, float64) float64, a, b float64) float64 {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with args "+"(%d, %d)", opName, a, b)
	return op(a, b)
}
func TestFunctional(t *testing.T) {
	res := apply(math.Pow, 2, 3)
	t.Log(res)
}

// 3. Variadic function
func sum(nums ...int) int {
	res := 0
	for num := range nums {
		res += num
	}
	return res
}
func Test(t *testing.T) {
	t.Log(sum(1, 2, 3, 4, 5))
}

// 4. Keep it simple
/*
   No function override, no default parameter values

*/

