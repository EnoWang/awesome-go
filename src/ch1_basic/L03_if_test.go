package basic

import (
	"io/ioutil"
	"testing"
)

// 0. Can do assignment in if clause. The scopes of assigned variables are in this if statement
func TestBound(t *testing.T) {
	const filename = "abc.txt"
	if contents, err := ioutil.ReadFile(filename); err != nil {
		t.Log(err)
	} else {
		t.Log(contents)
	}
}

// 1. switch
// Note: no need to add break statement
func TestSwitch(t *testing.T) {
	score := 40
	switch {
	case score < 0 || score > 100:
		t.Log("invalid score")
	case score < 60:
		t.Log("F")

	case score < 80:
		t.Log("C")

	case score < 90:
		t.Log("B")

	case score < 100:
		t.Log("A")
	}
}
