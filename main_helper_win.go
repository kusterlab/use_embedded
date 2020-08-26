// +build windows

package main

import "os/exec"
import "fmt"

func openURL(url string) error {
	fmt.Println("test")
	return exec.Command("cmd.exe", "/C", "start "+url).Run() // #nosec
}
