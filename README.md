This repo uses Go fuzzing in Go 1.18 to fuzz Dockerfile parsing in
https://pkg.go.dev/github.com/moby/buildkit/frontend/dockerfile/parser#Parse

It's mainly a testbed to play with fuzzing, but who knows, maybe it'll find
something interesting.
