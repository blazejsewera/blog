.PHONY: run clean test archive

help: build
	./renderblog -h

build:
	go build -o renderblog github.com/blazejsewera/blog

clean:
	rm -f renderblog

test:
	go test ./...

archive:
	tar -cvf ./dist.tar ./dist
