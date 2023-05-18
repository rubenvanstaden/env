SHELL := /bin/bash

UNIT_TEST_PATH=./...

fmt:
	GOPROXY="" go mod tidy -compat=1.17
	gofmt -l -s -w .

test:
	go test -count=1 -run=Unit $(UNIT_TEST_PATH)

debug:
	go test -count=1 -run=Unit $(UNIT_TEST_PATH) -v
