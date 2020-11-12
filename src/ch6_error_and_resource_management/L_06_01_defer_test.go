package ch6_error_and_resource_management

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

// 0. defer: make sure happen when function is finishing
func TestBasicDefer(t *testing.T) {
	defer t.Log(1)
	t.Log(2)
	t.Log(3)
}

// 1. multiple defer
// Note: multiple is like a stack, FILO
func TestMultipleDefer(t *testing.T) {
	defer t.Log(1)
	defer t.Log(2)
	t.Log(3)
}

// 2. defer works even there is return or panic
// Note: multiple is like a stack, FILO
func TestAlwaysDefer(t *testing.T) {
	defer t.Log("This is the test output for defer statement")
	panic("test panic")
}

// 3. defer for resource management
func Fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func writeFile(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// write in buffer
	writer := bufio.NewWriter(file)
	defer writer.Flush()

	f := Fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}
}

func TestWriteFile(t *testing.T) {
	writeFile("fib.txt")
}
