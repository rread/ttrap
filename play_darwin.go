package main

import "os/exec"

func play(f string) error {
	cmd := exec.Command("afplay", "-t", "0.1", f)
	return cmd.Run()

}
