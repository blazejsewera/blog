build:
	cd renderer; $(MAKE) build

test:
	cd renderer; $(MAKE) test

clean:
	rm -f ./renderblog

archive:
	tar -cvf ./dist.tar ./dist
