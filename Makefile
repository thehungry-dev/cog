.PHONY: clean build test

build:
	go build

test:
	go test -v

clean:
	go clean -i github.com/theungry-dev/log...
