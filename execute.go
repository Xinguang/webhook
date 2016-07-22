package main

import (
	"bytes"
	"log"
	"os/exec"
)

func executeShell(shell string) {

	if shell == "" {
		return
	}

	go func() {
		command := exec.Command("sh", "-c", shell)
		var out bytes.Buffer
		command.Stdout = &out
		command.Stderr = &out
		err := command.Run()

		if err != nil {
			log.Printf(out.String())
		} else {
			log.Printf("Shell output was: %s\n", out.String())
		}

	}()
}
