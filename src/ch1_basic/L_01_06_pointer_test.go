package basic

import "testing"

// 0. pointer
// We can not do operation on pointer like C
func TestPointer(t *testing.T) {
	a := 2
	var pa = &a
	*pa = 3
	t.Log(a)
}

// 1. only pass-by-value in go. So we need to leverage pointer to update variable in calling function
