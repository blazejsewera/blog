.PHONY: run clean test

run: build
	./renderblog

run-f: clean build
	./renderblog -f=1

run-ff: clean build
	./renderblog -f=2

build: renderblog

clean:
	rm -f renderblog

renderblog:
	go build -o renderblog github.com/blazejsewera/blog

test:
	go test ./...
