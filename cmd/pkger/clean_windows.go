// +build windows

package main

import (
	"os"
	"os/exec"
	"syscall"
)

func clean() {
	c := exec.Command("go", "mod", "tidy")
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	c.Stdin = os.Stdin
	c.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	c.Run()
}
