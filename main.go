package main

import (
	"runtime"
)

const VERSION = "0.1.0-alpha"

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	SetupWebServer()
}