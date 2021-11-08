package main

import (
	"os"
	"github.com/mrtc0/monkey/src/monkey/repl"
)

func main() {
	repl.Start(os.Stdin,os.Stdout)
}