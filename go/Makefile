.PHONY: run

run: build
	./renderblog

run-f: build
	./renderblog -f=1

run-ff: build
	./renderblog -f=2

build: renderblog

renderblog:
	go build -o renderblog github.com/blazejsewera/blog
