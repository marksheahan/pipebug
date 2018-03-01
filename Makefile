default: test

test: force
	go test -count 100

fmt: force
	gofmt -w *.go

.PHONY: default test force fmt

