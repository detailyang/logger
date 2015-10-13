GOPATH := ${GOPATH}:$(shell pwd)
.PHONY: clean test

all:
		@GOPATH=$(GOPATH) go build -ldflags "-X main.Buildstamp `date  '+%Y-%m-%d_%k:%M:%S'` -X main.Githash `git rev-parse HEAD`" ylogger
		rm -rf bin
		mkdir bin
		mv ylogger bin

clean:
		@rm -fr bin pkg

test:
		@GOPATH=$(GOPATH) go test logger
