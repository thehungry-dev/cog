.PHONY: clean build test bench

BENCHTIME ?= 1s
BENCH_FLAGS ?= -cpuprofile=cpu.pprof -memprofile=mem.pprof -benchmem -benchtime=$(BENCHTIME)

build:
	go build

test:
	go test -v ./test/automated/...

bench:
	cd benchmark; \
	go test -v -bench=. -run="^$$" $(BENCH_FLAGS);

clean:
	go clean -i github.com/theungry-dev/log...

interactive:
	go run test/interactive/$(TEST)/main.go
