// +build !windows

package main

import (
	"os/exec"
	"runtime"
	"syscall"
)

func openURL(url string) error {
	switch runtime.GOOS {
	case "darwin":
		return exec.Command("open", url).Run() // #nosec

	default:
		cmd := exec.Command("xdg-open", url) // #nosec
		cmd.SysProcAttr = &syscall.SysProcAttr{
			Setpgid: true,
		}
		return cmd.Run()
	}
}
