package main

import "os/exec"

func play(f string) error {
	cmd := exec.Command("afplay", f)
	return cmd.Run()

}
