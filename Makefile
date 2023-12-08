build:
	cd renderer; $(MAKE) build

test:
	cd renderer; $(MAKE) test

clean:
	rm -f ./renderblog

archive:
	tar -cvf ./dist.tar ./dist

dev: install build-render
	go run devserver/main.go &
	reflex --config=devserver.reflex --decoration=fancy

build-render: build
	./renderblog -f=1

install:
ifeq (, $(shell which reflex))
	go install github.com/cespare/reflex@latest
	@echo "Remember to add Go bin directory into PATH (usually $HOME/go/bin)"
endif
