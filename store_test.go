package main

import (
	"bytes"
	"fmt"
	"testing"
)

func TestPathTransformFunc(t *testing.T) {
	key := "bestpic"
	pathname := CASPathTransformFunc(key)
	fmt.Print(pathname)
	expected := "6f90c/0cbff/d1b2a/a1e69/c839a/5b960/6ff14/5c565"
	if pathname != expected {
		t.Errorf("Need %s Got %s", pathname, expected)
	}
}

func TestStore(t *testing.T) {
	opts := StoreOpts{
		PathTransformFunc: CASPathTransformFunc,
	}

	s := NewStore(opts)

	data := bytes.NewReader([]byte("some jpb bytes"))
	if err := s.writeStream("mypic", data); err != nil {
		t.Error(err)
	}
}
