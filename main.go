package main

import (
	"bufio"
	"log"
	"os"
	"strings"
	"tcpserver/kqueue"
	"tcpserver/socket"
)

func main() {
	s, err := socket.Listen("127.0.0.1", 8080)
	if err != nil {
		log.Println("Failed to create Socket:", err)
		os.Exit(1)
	}

	eventLoop, err := kqueue.NewEventLoop(s)
	if err != nil {
		log.Println("Failed to create event loop:", err)
		os.Exit(1)
	}
	log.Println("Server started. Waiting for incoming connections. ^C to exit.")

	eventLoop.Handle(func(socket *socket.Socket) {
		reader := bufio.NewReader(s)
		for {
			line, err := reader.ReadString('\n')
			if err != nil || strings.TrimSpace(line) == "" {
				break
			}
			_, err = s.Write([]byte(line))
			if err != nil {
				return
			}
		}
		err := s.Close()
		if err != nil {
			return
		}
	})

}
