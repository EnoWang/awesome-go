package ch2_built_in_collections

import (
	"fmt"
	"testing"
)

// 0. declare slice
// Note: splice is a view of array
func TestBasicSlice(t *testing.T) {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	t.Log("arr[2:6] =", arr[2:6])
	t.Log("arr[:6] =", arr[:6])
	t.Log("arr[2:] =", arr[2:])
	t.Log("arr[:] =", arr[:])
}

// 1. update original array
func updateSlice(s []int) {
	s[0] = 100
}
func TestUpdate(t *testing.T) {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := arr[2:]
	t.Log("s1 =", s1)
	s2 := arr[:]
	t.Log("s2 =", s2)

	t.Log("After updateSlice(s1)")
	updateSlice(s1)
	t.Log("s1 =", s1)
	t.Log(arr)
	t.Log("s2 =", s2)
}

// 2. reslice
func TestReslice(t *testing.T) {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := arr[2:]
	t.Log("s1 =", s1)
	s2 := s1[2:]
	t.Log("s2 =", s2)
}

// 3. extendable
// Note: slice has len and cap, so if the upper bound index < cap, there
// is no problem, else there will be an error
func TestExtend(t *testing.T) {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	s1 := arr[2:6]
	t.Log("s1 =", s1)
	s2 := s1[3:5]
	t.Log("s2 =", s2)
	// Note: arr
	//s3 := s1[3:7]
	//t.Log("s3 =", s3)
	t.Logf("s1=%v, len(s1)=%d, cap(s1)=%d\n", s1, len(s1), cap(s1))
}

// 4. add element to slice
// note: When adding element in slice, if more elements than cap, there will be a new array
// which contains all element copied
// Note: The old array will be garbage collected
// Note: Need to receive return value of append
func TestAddElement(t *testing.T) {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	s1 := arr[2:6]
	s2 := s1[3:5]
	s3 := append(s2, 10)
	s4 := append(s3, 11)
	s5 := append(s4, 12)
	t.Logf("s3=%v, s4=%v, s5=%v\n", s3, s4, s5)
	t.Log("arr\n", arr)
}

// 5. create slice
// 5.1 Create slice
func printSlice(s []int) {
	fmt.Printf("len=%d, cap=%d\n", len(s), cap(s))
}
func TestCreate(t *testing.T) {
	// Note: zero value for slice is nil
	var s []int
	// Note: Will always skip bc s is nil
	for _, v := range s {
		t.Log(v)
	}

	// Note:
	for i := 0; i < 100; i++ {
		s = append(s, 2*i+1)
		printSlice(s)
	}
}

// 5.2 Create slice with initial elements
func TestCreateWithVals(t *testing.T) {
	s := []int{2, 4, 6, 8}
	printSlice(s)
}

// 5.3 Create slice bt built-in make function
// Note: can specify capacity
func TestCreateByMake(t *testing.T) {
	s1 := make([]int, 16)
	printSlice(s1)
	t.Log(s1)

	s2 := make([]int, 16, 32)
	printSlice(s2)
	t.Log(s2)
}

// 6. copy slice
// Note: built_in function copy(dest, source)
func TestCopy(t *testing.T) {
	s1 := []int{2, 4, 6, 8}
	s2 := make([]int, 16)
	copy(s2, s1)
	t.Log(s2)
}

// 7. delete element in slice
// use append
func TestDeleteSlice(t *testing.T) {
	t.Log("Original array")
	s := []int{2, 4, 6, 8}
	t.Log(s)

	t.Log("Delete element in the middle of slice")
	s1 := append(s[:2], s[3:]...)
	t.Log(s1)

	t.Log("Pop from front")
	s2 := s[1:]
	t.Log(s2)

	t.Log("Pop from tail")
	s3 := s[:len(s)-1]
	t.Log(s3)
}
