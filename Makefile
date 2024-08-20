all: run_main

build: guest.wasm

# NOTE: Install `gotip` first, see how https://pkg.go.dev/golang.org/dl/gotip
.PHONY: guest.wasm
guest.wasm:
	GOOS=wasip1 GOARCH=wasm gotip build -buildmode=c-shared -o ./guest/guest.wasm ./guest

# NOTE: Install `wasmtime` first, see how https://docs.wasmtime.dev/cli-install.html#installing-wasmtime
.PHONY: run_wasmtime_cli
run_wasmtime_cli:
	wasmtime --invoke CallAdd ./guest/guest.wasm

.PHONY: run_main
run_main:
	go run .
