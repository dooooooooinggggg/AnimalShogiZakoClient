package main

import (
	"bufio"
	"fmt"
	"log"
	"net"

	"github.com/dooooooooinggggg/AnimalShogiZakoClient/connect"
)

func main() {
	conn := connect.TcpConnection()
	defer conn.Close()
	fmt.Printf("Connected to 'Animal Shogi Server'\n")
	sendMessage(conn)
}

func sendMessage(conn net.Conn) {
	text := "BEGIN Game_Summary"
	log.Print("Will send text: " + text)
	conn.Write([]byte(text + "\n"))
	message, _ := bufio.NewReader(conn).ReadString('\n')
	log.Print("Message from server: " + message)
}
