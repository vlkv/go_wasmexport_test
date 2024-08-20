package wasmtime_helper

import (
	"io"
	"os"

	wasmtime "github.com/bytecodealliance/wasmtime-go/v23"
)

func NewWasmInstance(wasmBinaryPath string) (*wasmtime.Store, *wasmtime.Instance, error) {
	engine := wasmtime.NewEngine()

	linker := wasmtime.NewLinker(engine)
	err := linker.DefineWasi()
	check(err)

	store := wasmtime.NewStore(engine)

	wasiConfig := wasmtime.NewWasiConfig()
	wasiConfig.InheritArgv()
	wasiConfig.InheritEnv()
	wasiConfig.InheritStdin()
	wasiConfig.InheritStderr()
	wasiConfig.InheritStdout()
	wasiConfig.PreopenDir(".", "/")
	store.SetWasi(wasiConfig)

	wasmBytes, err := ReadWasmBytes(engine, wasmBinaryPath)
	check(err)

	module, err := wasmtime.NewModule(engine, wasmBytes)
	check(err)

	// instance, err := wasmtime.NewInstance(store, module, []wasmtime.AsExtern{})
	instance, err := linker.Instantiate(store, module)
	check(err)

	return store, instance, nil
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadWasmBytes(engine *wasmtime.Engine, path string) ([]byte, error) {
	wasmFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	wasm, err := io.ReadAll(wasmFile)
	if err != nil {
		return nil, err
	}

	if err := wasmtime.ModuleValidate(engine, wasm); err != nil {
		return nil, err
	}

	return wasm, nil
}
