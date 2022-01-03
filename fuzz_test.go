package fuzz_test

import (
	"bytes"
	"testing"

	"github.com/moby/buildkit/frontend/dockerfile/parser"
)

func FuzzDockerfileParse(f *testing.F) {
	f.Add([]byte(`FROM ubuntu`))

	f.Fuzz(func(t *testing.T, content []byte) {
		if _, err := parser.Parse(bytes.NewReader(content)); err != nil {
			t.Logf("%s: %v", string(content), err)
		}
	})
}
