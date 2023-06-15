package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("docker", "build", "-t", "job:latest", ".")
	// cmd := exec.Command("echo", "build", "-t job:latest", ".")
	stderr, _ := cmd.StderrPipe()
	defer stderr.Close()
	// output, _ := cmd.StdoutPipe()

	wait := make(chan bool)

	go func() {
		reader := bufio.NewReader(stderr)
		buf := make([]byte, 0, 24)
		for {
			n, err := io.ReadFull(reader, buf[:cap(buf)])
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
			fmt.Print(string(buf))
		}

		wait <- true
	}()

	go func() {
		// scanner := bufio.NewScanner(stderr)
		// for scanner.Scan() {
		// fmt.Print(scanner.Text())
		// }
		wait <- true
	}()

	err := cmd.Start()
	if err != nil {
		fmt.Println(err.Error())
	}

	err = cmd.Wait()
	<-wait
	<-wait

	if err != nil {
		fmt.Println(err.Error())
	}

}
