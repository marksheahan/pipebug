default: test

test: force
	go test -count 100


.PHONY: default test force
