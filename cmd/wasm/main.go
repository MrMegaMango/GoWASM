//go:build js && wasm
// +build js,wasm

package main

import (
	"fmt"
	"syscall/js"
	"time"
)

func main() {
	fmt.Println("hello WASM")
	js.Global().Set("SaySomething", SaySomething())
	for {
		time.Sleep(3 * time.Second)
	}
}

func SaySomething() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		fmt.Println("something woo woo")
		return `<h4>Go logic is called</h4>`
	})
}
