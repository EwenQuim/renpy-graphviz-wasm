package main

import (
	"fmt"
	"strings"
	"syscall/js"

	"pkg.amethysts.studio/renpy-graphviz/parser"
)

func main() {
	done := make(chan bool)

	js.Global().Set("printMessage", js.FuncOf(printMessage))

	fmt.Println("exiting")
	<-done
	fmt.Scanln()
	fmt.Println("exited")

}

func printMessage(this js.Value, inputs []js.Value) interface{} {
	callback := inputs[len(inputs)-1:][0]

	fmt.Println("input", inputs[0].String())

	// api.github.com/search/code?accept=application/vnd.github.v3+json&q=repo:amethysts-studio/coalescence+extension:rpy
	renpyRepoCodeLines := strings.Split(inputs[0].String(), "\n") //[]string{"label hello:", "world", "jump label2"} //getRenpyFromRepo(inputs[0].String())

	fmt.Println("string inside Go - renpy lines", renpyRepoCodeLines)

	dotGraph := parser.Graph(renpyRepoCodeLines)

	fmt.Println("string inside Go - graph", dotGraph.String())

	// document := js.Global().Get("document")
	// p := document.Call("createElement", "p")
	// p.Set("innerHTML", dotGraph.String())
	// document.Get("body").Call("appendChild", p)
	callback.Invoke(js.Null(), dotGraph.String())
	return nil
}
