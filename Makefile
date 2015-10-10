GOPATH := ${GOPATH}:$(shell pwd)
.PHONY: clean test

all:
		@GOPATH=$(GOPATH) go install logger


clean:
		@rm -fr bin pkg

test:
		@GOPATH=$(GOPATH) go test logger
