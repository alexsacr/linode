.PHONY: default test test-var integration record clean cover coverhtml test-echo setup

default: test

test: clean
	go test -cover
	go vet
	golint
	errcheck

test-all:
	go test -v
	go test -v -tags="integration" -timeout 20m

test-var: clean
ifeq ($(TEST),)
	$(error TEST not set.)
endif

integration: clean test-var
	go test -v -tags="integration" -run=$(TEST) -timeout 15m

record: clean test-var
	go test -v -tags="integration debug" -run=$(TEST)

clean:
	rm -f recorded.go

cover:
	go test -coverprofile=coverage.out

coverhtml: cover
	go tool cover -html=coverage.out -o coverage.html

test-echo:
	go test -v -tags="integration" -run=TestTestEcho

setup:
	go get github.com/golang/lint/golint
	go get github.com/kisielk/errcheck