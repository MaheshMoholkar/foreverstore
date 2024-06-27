package main

import (
	"bytes"
	"io"
	"testing"
)

func TestPathTransformFunc(t *testing.T) {
	key := "bestpic"
	pathKey := CASPathTransformFunc(key)
	expectedOriginal := "6f90c0cbffd1b2aa1e69c839a5b9606ff145c565"
	expectedPathname := "6f90c/0cbff/d1b2a/a1e69/c839a/5b960/6ff14/5c565"

	if pathKey.Pathname != expectedPathname {
		t.Errorf("Need %s Got %s", pathKey.Pathname, expectedPathname)
	}

	if pathKey.Filename != expectedOriginal {
		t.Errorf("Need %s Got %s", pathKey.Filename, expectedOriginal)
	}
}

func TestStore(t *testing.T) {
	opts := StoreOpts{
		PathTransformFunc: CASPathTransformFunc,
	}

	s := NewStore(opts)
	key := "specialkey"

	data := []byte("some jpg bytes")
	if err := s.writeStream(key, bytes.NewReader(data)); err != nil {
		t.Error(err)
	}

	r, err := s.Read(key)
	if err != nil {
		t.Error(err)
	}
	b, _ := io.ReadAll(r)
	if string(b) != string(data) {
		t.Errorf("Need %s but Got %s", data, b)
	}
}
