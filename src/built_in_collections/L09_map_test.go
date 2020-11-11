package built_in_collections

import "testing"

// 0. declare map
// Note: we can do operations on nil in golang
func TestBasicMap(t *testing.T) {
	m1 := map[string]string{
		"name": "golang",
		"ide":  "goland",
	}
	t.Log(m1)

	m2 := make(map[string]int) // m2 == empty map

	var m3 map[string]int // m3 == nil

	t.Log(m2)
	t.Log(m3)
}

// 1. Traverse map
// Note: we can do operations on nil in golang
// Note: no order in map
func TestTraverse(t *testing.T) {
	m := map[string]string{
		"name": "golang",
		"ide":  "goland",
	}
	for k, v := range m {
		t.Log(k, v)
	}
}

// 2. get value
// Note: print out zero value for invalid key
func TestValue(t *testing.T) {
	m := map[string]string{
		"name": "golang",
		"ide":  "goland",
	}
	t.Log(m["name"])
	// Note: zero value for string is empty string
	t.Log(m["invalid"])
}

// 3. check contains
func TestContains(t *testing.T) {
	m := map[string]string{
		"name": "golang",
		"ide":  "goland",
	}
	if name, ok := m["name"]; ok {
		t.Log(name)
	} else {
		t.Log("invalid key")
	}
}

// 4. delete element
func TestDeleteMap(t *testing.T) {
	m := map[string]string{
		"name": "golang",
		"ide":  "goland",
	}
	if name, ok := m["name"]; ok {
		t.Log(name)
	} else {
		t.Log("invalid key")
	}

	delete(m, "name")

	if name, ok := m["name"]; ok {
		t.Log(name)
	} else {
		t.Log("invalid key")
	}
}

// 5 type of key
// All built_in types can be key excluding map, slice and function
// Struct without fields in function, map or function can be key


// 6. longest substring without repeating characters
func getLongestSubstringWithoutRepeatingChars(s string) int {
	lastOccurred := make(map[byte]int)
	start := 0
	maxLength := 0
	for i, ch := range []byte(s) {

		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastOccurred[ch] + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxLength
}

func TestLS(t *testing.T) {
	t.Logf("s = %s, longest substirng with no repeating characters: %d\n", "abcabcabc", getLongestSubstringWithoutRepeatingChars("abcabcabc"))
}
