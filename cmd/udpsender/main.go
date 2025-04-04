package main

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"

	// "log/slog"
	"net"
)

func main() {
	addr, err := net.ResolveUDPAddr("udp", "127.0.0.1:8080")
	if err != nil {
		panic(err)
	}
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	rd := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		str, err := rd.ReadString('\n')
		if err != nil {
			slog.Error("Error reading string", "error", err)
			fmt.Println(err)
			continue
		}

		_, err = conn.Write([]byte(str))
		if err != nil {
			slog.Error("Error writing string", "error", err)
			continue
		}
	}
}
