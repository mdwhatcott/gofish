#!/usr/bin/make -f

test:
	go build ./...
	go test ./...
