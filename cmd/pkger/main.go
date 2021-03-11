package main

import (
	"os"

	"github.com/ambitvpn/pkger/cmd/pkger/cmds"
)

func main() {
	root, err := cmds.New()
	if err != nil {
		root.Route(os.Args[1:])
	}
}
