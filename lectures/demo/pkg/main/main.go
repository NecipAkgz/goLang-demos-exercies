// run with `go run ./demo/pkg/main`
package main

import (
	"goLearning/lectures/demo/pkg/display"
	"goLearning/lectures/demo/pkg/msg"
)

func main() {
	msg.Hi()
	display.Display("hello from display")
	msg.Exciting("calling a package function")
}
