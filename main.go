package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

const FILE_NAME = "messages.txt"

func main() {
	file, err := os.Open(FILE_NAME)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var line = ""

	for {
		buf := make([]byte, 8)
		n, err := file.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		if n == 0 {
			break
		}

		parts := strings.Split(string(buf), "\n")
		if len(parts) > 1 {
			line = fmt.Sprintf("%s%s", line, parts[0])
			fmt.Println("read:", line)
			line = parts[1]
			continue
		}

		line = fmt.Sprintf("%s%s", line, string(buf))
	}
}
