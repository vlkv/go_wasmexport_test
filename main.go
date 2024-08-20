package main

import wasm "go_wasmexport_test/wasmtime_helper"

func main() {
	store, instance, err := wasm.NewWasmInstance("./guest/guest.wasm")
	if err != nil {
		panic(err)
	}

	mainFunc := instance.GetFunc(store, "_start")
	_, err = mainFunc.Call(store)
	if err != nil {
		panic(err)
	}

	addFunc := instance.GetFunc(store, "Add")
	_, err = addFunc.Call(store, 40, 2)
	if err != nil {
		panic(err)
	}
}
