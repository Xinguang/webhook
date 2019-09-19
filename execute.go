package main

import (
	"bufio"
	"context"
	"os/exec"
	"time"

	log "github.com/sirupsen/logrus"
)

func executeShell(shell string) {

	if shell == "" {
		return
	}

	go execute(shell)
}

func execute(shell string) {
	ctx, calcel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer calcel()

	cmd := exec.CommandContext(ctx, "sh", "-c", shell)

	stdout, _ := cmd.StdoutPipe()
	stderr, _ := cmd.StderrPipe()
	oReader := bufio.NewReader(stdout)
	eReader := bufio.NewReader(stderr)
	scan(oReader, log.Info)
	scan(eReader, log.Info)

	err := cmd.Start()
	if err != nil {
		log.Error(err)
	}

	err = cmd.Wait()
	if err != nil {
		log.Error(err)
	}
}

func scan(reader *bufio.Reader, f func(args ...interface{})) {
	go func() {
		for {
			line, err := reader.ReadString('\n')
			if len(line) > 0 {
				f(line)
			}
			if err != nil {
				break
			}
		}
	}()
}
