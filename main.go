package main

import (
	"fmt"
	"io"
	"net"
	"strings"
)

const FILE_NAME = "messages.txt"

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}

		fmt.Println("client connected")

		ch := getLinesChannel(conn)
		for line := range ch {
			fmt.Println(line)
		}
		fmt.Println("client disconnected")
	}

}

func getLinesChannel(f io.ReadCloser) <-chan string {
	var line = ""
	var ch = make(chan string)

	go func() {
		for {
			buf := make([]byte, 8)
			n, err := f.Read(buf)
			if n == 0 {
				// NOTE: The order is important when n == 0 err == io.EOF
				break
			}
			if err != nil {
				panic(err)
			}

			parts := strings.Split(string(buf), "\n")
			if len(parts) > 1 {
				line = fmt.Sprintf("%s%s", line, parts[0])
				ch <- line
				line = parts[1]
				continue
			}

			line = fmt.Sprintf("%s%s", line, string(buf))
		}

		// NOTE: Remmember that a channel must always be closed
		close(ch)
	}()

	return ch
}
