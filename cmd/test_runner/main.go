package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"time"
)

func main() {
	cmd := exec.Command("docker",
		"build",
		"-t",
		"job:latest",
		".",
	)
	// cmd := exec.Command("echo", "build", "-t job:latest", ".")
	stderr, _ := cmd.StderrPipe()
	defer stderr.Close()
	stdout, _ := cmd.StdoutPipe()
	defer stdout.Close()

	wait := make(chan bool)

	go func() {
		merged := io.MultiReader(stderr, stdout)
		reader := bufio.NewReader(merged)
		for {
			buf := make([]byte, 32)
			n, err := reader.Read(buf)
			buf = buf[:n]

			if err != nil {
				if err == io.EOF {
					break
				}

				if err != io.ErrUnexpectedEOF {
					fmt.Fprintln(os.Stderr, err)
					break
				}
			}

			_ = time.Second * 2
			// time.Sleep(time.Millisecond * 10)
			fmt.Print(string(buf))
		}

		wait <- true
	}()

	err := cmd.Start()
	if err != nil {
		fmt.Println(err.Error())
	}

	<-wait
	err = cmd.Wait()

	if err != nil {
		fmt.Println(err.Error())
	}

}
