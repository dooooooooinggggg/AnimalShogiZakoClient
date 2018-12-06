package main

import (
	"bufio"
	"fmt"
	"log"
	"net"

	"github.com/dooooooooinggggg/AnimalShogiZakoClient/connect"
)

func main() {
	conn := connect.TCPConnection()
	defer conn.Close()
	fmt.Printf("Connected to 'Animal Shogi Server'\n")
	go connect.KeepAlive(conn)
	sendMessage(conn)
}

func sendMessage(conn net.Conn) {
	for i := 0; i < 10; i++ {
		conn.Write([]byte("\n"))
		message, _ := bufio.NewReader(conn).ReadString('\n')
		log.Print("Message from server: " + message)
		break
	}
}
