package ch4_interface_oriented

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"testing"
	"time"
)

// 0. Interface implementation
type Retriever interface {
	Get(url string) (*string, error)
}

type mockRetriever struct {
	Contents string
}

func (r mockRetriever) Get(url string) (*string, error) {
	return &(r.Contents), nil
}

type retriever struct {
	UserAgent string
	TimeOut   time.Duration
}
func (r *retriever) Get(url string) (*string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	result, err := httputil.DumpResponse(resp, true)

	resp.Body.Close()
	if err != nil {
		return nil, err
	}
	res := string(result)
	return &res, nil
}

func download(r Retriever) string {
	if response, err := r.Get("http://www.google.com"); err != nil {
		panic(err)
	} else {
		return *response
	}
}

func TestInterBasicInterface(t *testing.T) {
	r := &retriever{}
	t.Log(download(r))
}

// 1. Digging into implementations for Retriever
func inspect(r Retriever) {
	fmt.Printf("%T %v\n", r, r)
	switch v := r.(type) {
	// Note: use switch to check the type of interface
	case mockRetriever:
		fmt.Println("Contents:", v.Contents)
	case *retriever:
		fmt.Println("UserAgent:", v.UserAgent)
	}
}

func TestImplementation(t *testing.T) {
	r1 := mockRetriever{
		Contents: "Hello World",
	}
	fmt.Println("----------------- inspection for r1 -----------------------")
	inspect(r1)
	r2 := &retriever{
		UserAgent: "mock",
		TimeOut:   time.Minute,
	}
	fmt.Println("----------------- inspection for r2 -----------------------")
	inspect(r2)
}

// 2. Type Assertion: check the type of interface
func TestTypeAssertion(t *testing.T) {
	var r Retriever
	r = &retriever{
		UserAgent: "mock",
		TimeOut:   time.Minute,
	}

	// wrong way!
	//t.Log(r.TimeOut)
	if rt, ok := r.(*retriever); ok {
		t.Log(rt.TimeOut)
	} else {
		t.Log("Not a *retriever")
	}
}

// 3. interface{}: any type of interface
type Queue []interface{}

// Note: queue will be able to accept any type of elements
/*
func (q *Queue) Push(v interface{}) {
	*q = append(*q, v)
}

func (q *Queue) Poll() interface{} {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}
*/

func (q *Queue) Push(v int) {
	*q = append(*q, v)
}

func (q *Queue) Poll() int {
	head := (*q)[0]
	*q = (*q)[1:]
	return head.(int)
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}

func TestQueue(t *testing.T) {
	q := Queue{1}
	q.Push(2)
	q.Push(3)
	t.Log(q.Poll())
	t.Log(q.Poll())
	t.Log(q.IsEmpty())
	t.Log(q.Poll())
	t.Log(q.IsEmpty())
}
