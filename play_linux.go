package main

import "os/exec"

func play(f string) error {
	cmd := exec.Command("aplay", f)
	return cmd.Run()

}
