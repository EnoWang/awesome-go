package ch4_interface_oriented

import (
	"testing"
)

type Retriever interface {
	Get(url string) (*string, error)
}

type Poster interface {
	Post(url string, form map[string]string) string
}

type RetrieverPoster interface {
	Retriever
	Poster
}

type mockRetrieverPoster struct {
	Contents string
}

func (rp *mockRetrieverPoster) Get(url string) (*string, error) {
	return &(rp.Contents), nil
}

func (rp *mockRetrieverPoster) Post(url string, form map[string]string) string {
	rp.Contents = form["contents"]
	return "ok"
}

func post(poster Poster) {
	poster.Post("http://www.google.com", map[string]string{
		"name": "golang",
		"ide":  "goland",
	})
}

func session(rp RetrieverPoster) string {
	rp.Post("http://www.google.com", map[string]string{
		"contents": "test contents",
	})
	resp, err := rp.Get("http://www.google.com")
	if err != nil {
		return "error"
	} else {
		return *resp
	}
}

func TestInterfaceCombination(t *testing.T) {
	r := &mockRetrieverPoster{
		Contents: "Hello World",
	}
	t.Log(session(r))
}
