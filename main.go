package main

import (
	"bytes"
	"fmt"
	"io"
	"log/slog"
	"os"
)

func main() {
	file, err := os.Open("messages.txt")

	if err != nil {
		slog.Error("Unable to open the file", "err", err)
		os.Exit(1)
	}
	defer func() {
		err := file.Close()
		if err != nil {
			slog.Error("Unable to close the file", "err", err)
			os.Exit(1)

		}
	}()

	str := ""
	for {
		buf := make([]byte, 8)
		n, err := file.Read(buf)
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
			fmt.Printf("read: %s\n", str)
			str = ""
		}
		str += string(buf)

	}

	if len(str) > 0 {
		fmt.Printf("read: %s\n", str)
	}
}
