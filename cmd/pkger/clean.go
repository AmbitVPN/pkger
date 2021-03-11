// +build !windows

package main

import (
	"os"
	"os/exec"
)

func clean() {
	c := exec.Command("go", "mod", "tidy")
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	c.Stdin = os.Stdin
	c.Run()
}
