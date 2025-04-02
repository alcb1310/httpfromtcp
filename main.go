package main

import (
	"fmt"
	"io"
	"os"
)

const FILE_NAME = "messages.txt"

func main() {
	file, err := os.Open(FILE_NAME)
	if err != nil {
		panic(err)
	}
	defer file.Close()

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
		fmt.Println("read:", string(buf))
	}
}
