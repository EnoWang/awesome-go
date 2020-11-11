package _2_built_in_collections

import (
	"fmt"
	"testing"
)

// 0. declare array
func TestBasicArray(t *testing.T) {
	// Note: will have default values
	var arr1 [5]int
	arr2 := [3]int{1, 3, 5}
	t.Log(arr1, arr2)
}

// 1. declare array with ...
// Note: compiler can inference size
func TestWithoutSize(t *testing.T) {
	arr := [...]int{2, 4, 6, 8, 10}
	t.Log(arr)
}

// 2. multiple dimension
func TestMutiDimension(t *testing.T) {
	var grid [4][5]bool
	t.Log(grid)
}

// 3. iterate array
func TestIterate(t *testing.T) {
	arr := [...]int{2, 4, 6, 8, 10}
	for i, v := range arr {
		t.Log(i, v)
	}
	for _, v := range arr {
		t.Log(v)
	}
}

// 4, array is value not reference
func printArr(arr [5]int) {
	arr[0] = 100
	for _, v := range arr {
		fmt.Println(v)
	}
}
func printArrByPtr(arr *[5]int) {
	// Note: can modify the value itself
	arr[0] = 100
	for _, v := range arr {
		fmt.Println(v)
	}
}
func TestArrType(t *testing.T) {
	arr := [...]int{2, 4, 6, 8, 10}
	// Note: these two are diff
	printArr(arr)
	fmt.Println(arr)

	// Note: these two are the same
	printArrByPtr(&arr)
	fmt.Println(arr)
}
