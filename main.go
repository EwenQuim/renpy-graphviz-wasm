package main

import (
	"fmt"
	"syscall/js"
	"time"

	"github.com/ewenquim/renpy-graphviz/parser"
)

func main() {
	start := time.Now()
	// api.github.com/search/code?accept=application/vnd.github.v3+json&q=repo:amethysts-studio/coalescence+extension:rpy
	renpyRepoCodeLines := getRenpyFromRepo("amethysts-studio/coalescence")

	fmt.Println("fetch", time.Since(start))

	dotGraph := parser.Graph(renpyRepoCodeLines)

	fmt.Println("parse", time.Since(start))

	fmt.Println(dotGraph.String())

	document := js.Global().Get("document")
	p := document.Call("createElement", "p")
	p.Set("innerHTML", dotGraph.String())
	document.Get("body").Call("appendChild", p)

}

// Runningtime computes running time
func runningtime(s string) (string, time.Time) {
	return s, time.Now()
}

// Track is this
func track(s string, startTime time.Time) {
	endTime := time.Now()
	fmt.Println(s, "took", endTime.Sub(startTime))
}
