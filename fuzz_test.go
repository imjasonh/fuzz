package fuzz_test

import (
	"bytes"
	"testing"

	"github.com/moby/buildkit/frontend/dockerfile/parser"
	"k8s.io/apimachinery/pkg/labels"
)

func FuzzDockerfileParse(f *testing.F) {
	f.Add([]byte(`FROM ubuntu`))

	f.Fuzz(func(t *testing.T, content []byte) {
		_, _ = parser.Parse(bytes.NewReader(content))
	})
}

func FuzzK8sLabelsParse(f *testing.F) {
	f.Add([]byte(`environment in (production, qa)`))

	f.Fuzz(func(t *testing.T, content []byte) {
		_, _ = labels.Parse(string(content))
	})
}
