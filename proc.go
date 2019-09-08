package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
)

var (
	isRunning = true
	sigChan = make(chan os.Signal, 1)
)

func StartProc(binaryPath string) {
	var err error
	var line = new([]byte)

	signal.Notify(sigChan, os.Interrupt)

	go func() {
		<-sigChan
		isRunning = false
	}()

	cmd := exec.Command(binaryPath)
	stdout, err := cmd.StdoutPipe()

	if err != nil {
		log.Fatal("cmd.StdoutPipe err:", err)
	}

	if err := cmd.Start(); err != nil {
		log.Fatal("cmd.Start err:", err)
	}

	stdoutReader := bufio.NewReader(stdout)

	for isRunning {
		*line, _, err = stdoutReader.ReadLine()
		if err != nil {
			log.Fatal("stdoutReader.ReadLine err:", err)
		}

		ParseLine(line)
	}

	fmt.Println("Exiting...")
	
	if err := cmd.Wait(); err != nil {
		log.Fatal("cmd.Wait err:", err)
	}
}
