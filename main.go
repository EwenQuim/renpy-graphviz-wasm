package main

import (
	"fmt"
	"syscall/js"

	"github.com/ewenquim/renpy-graphviz/parser"
)

func main() {

	// api.github.com/search/code?accept=application/vnd.github.v3+json&q=repo:amethysts-studio/coalescence+extension:rpy
	renpyRepoCodeLines := getRenpyFromRepo("amethysts-studio/coalescence")

	dotGraph := parser.Graph(renpyRepoCodeLines)

	fmt.Println(dotGraph.String())
	fmt.Println("q")

	document := js.Global().Get("document")
	p := document.Call("createElement", "p")
	p.Set("innerHTML", dotGraph)
	document.Get("body").Call("appendChild", p)

}
