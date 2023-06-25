package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func main() {
	pipe1 := strings.NewReader("billy")

	reader := bufio.NewReader(pipe1)
	// var result []byte

	for {
		buf := make([]byte, 3)
		n, err := reader.Read(buf)

		if err != nil {
			if err == io.EOF {
				fmt.Println("EOF")
			} else {
				fmt.Println(err.Error())
			}
			break
		}

		fmt.Println(n, string(buf[:n]))
	}
}
