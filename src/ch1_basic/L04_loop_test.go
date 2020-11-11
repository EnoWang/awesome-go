package basic

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"testing"
)

// 0. can skip start statement
func TestConvertToBin (t *testing.T) {
	n := 13
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	t.Log(result)
}

// 1. can skip start statement and increase statement
func TestReadFile (t *testing.T) {
	filename := "abc.txt"
	file, err := os.Open(filename)
	if  err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		t.Log(scanner.Text())
	}
}

// 2. can skip all statements
// Note: GoLang make it easy to write endless loop since goroutine need it
func TestForever (t *testing.T) {
	for {
		fmt.Println("Hello")
	}
}

// 3. no while in golang
