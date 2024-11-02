package main

import "syscall/js"

var x int = 0

func callback(this js.Value, args []js.Value) interface{} {
	js.Global().Call("alert", "Hello from Go-WebAssembly!")

	x++
	js.Global().Get("document").Call("getElementById", "result").Set("innerHTML", x)

	return js.ValueOf(0)
}

func main() {
	println("Hello from Go!")

	js.Global().Set("callback", js.FuncOf(callback))

	<-make(chan struct{}, 0)
}