package main

import (
	"log"
	"os"

	"github.com/ambitvpn/pkger/cmd/pkger/cmds"
)

func main() {
	defer clean()

	defer func() {
		if err := recover(); err != nil {
			clean()
			log.Fatal(err)
		}
	}()

	if err := run(); err != nil {
		clean()
		log.Fatal(err)
	}
}

func run() error {
	root, err := cmds.New()
	if err != nil {
		return err
	}

	return root.Route(os.Args[1:])
}

// does not compute
