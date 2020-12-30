package main

import (
	"github.com/dreblang/core/cmd/dreblc"
	// Core Lib
	_ "github.com/dreblang/corelib/http"
	_ "github.com/dreblang/corelib/math"
)

func main() {
	dreblc.Main()
}
