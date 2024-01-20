package main

import (
	"fmt"
	"syscall/js"
)

//export add
func add(a, b int) int {
	var c = a + b
	return c

}

func main() {
	fmt.Println("Hello World")
	js.Global().Set("goAdd", js.FuncOf(func(this js.Value, args []js.Value) any {
		return (add(args[0].Int(), args[1].Int()))

	}))
	select {}
}
