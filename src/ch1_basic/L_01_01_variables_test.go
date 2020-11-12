package basic

import (
	"math"
	"testing"
)

// 0. variable declared with default value
// Note: when printing empty string, use a combination of Printf and %q (as quote)
func TestDefaultValue(t *testing.T) {
	var a int
	var s string
	t.Logf("%d, %q\n", a, s)
}

// 1. Declare variables with initial values
// Note: can declare multiple variables in one statement
// Note: no values needed
func TestInitialValue(t *testing.T) {
	var a, b = 3, 4
	var s = "abc"
	t.Log(a, b, s)
}

// 2. Declare variables with initial values
// Note: in this way, variables in different types can declared in one statement
func TestTypeDeduction(t *testing.T) {
	var a, b, c, s = 3, 4, true, "def"
	t.Log(a, b, c, s)
}

// 3. Declare variables in short
// Note: assignment can not use := again
func TestShortDeclaration(t *testing.T) {
	a, b, c, s := 3, 4, true, "def"
	b = 5
	t.Log(a, b, c, s)
}

// 4. Declare variables out of functions
// Note: always need var key word
// Note: scope: in package
var (
	k   = "kk"
	bol = true
)

func TestPackageScopeVal(t *testing.T) {
	t.Log(k, bol)
}

// 5. built-in data types
/*
- bool, string
- (u)int, (u)int8, (u)int16, (u)int32, (u)int64, uintptr
- byte, rune
- float32, float64, complex64, complex128
*/

// 6. data type conversion: need to do it explicitly
func TestTriangle(t *testing.T) {
	var a, b = 3, 4
	var c = int(math.Sqrt(float64(a*a + b*b)))
	t.Log("triangle:", c)
}
