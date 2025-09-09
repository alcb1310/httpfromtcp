package main

import (
	"bytes"
	"fmt"
	"io"
	"log/slog"
	"net"
	"os"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")

	if err != nil {
		slog.Error("Unable to open the file", "err", err)
		os.Exit(1)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			slog.Error("Unable to accept the connection", "err", err)
			os.Exit(1)
		}

		for line := range getLinesChannel(conn) {
			fmt.Printf("read %s\n", line)
		}
	}
}

func getLinesChannel(f io.ReadCloser) <-chan string {
	out := make(chan string)

	go func() {
		defer f.Close()
		defer close(out)

		str := ""
		for {
			buf := make([]byte, 8)
			n, err := f.Read(buf)
			if err != nil {
				if err == io.EOF {
					break
				}

				slog.Error("Unable to read the file", "err", err)
				os.Exit(1)
			}

			buf = buf[:n]
			if idx := bytes.IndexByte(buf, '\n'); idx != -1 {
				str += string(buf[:idx])
				buf = buf[idx+1:]
				out <- str
				str = ""
			}
			str += string(buf)

		}

		if len(str) > 0 {
			out <- str
		}
	}()

	return out
}
