package main

import (
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

	for {
		buf := make([]byte, 8)
		n, err := file.Read(buf)
		if err != nil {
			if err == io.EOF {
				os.Exit(0)
			}

			slog.Error("Unable to read the file", "err", err)
			os.Exit(1)
		}

		fmt.Printf("read: %s\n", string(buf[:n]))
	}
}
