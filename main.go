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

func initConn(conn net.Conn) {
	conn.Write([]byte("LOGIN:Saikyou_AI\n"))
	conn.Write([]byte("AGREE\n"))
	message, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Printf("%s\n", message)
}

func sendMessage(conn net.Conn) {
	// ここ(対局相手が見つかったタイミング)で，自分が先手か後手か判定したい
	// てかここの初期化の処理書かないと
	// それまでここでストップしないと

	for {
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Printf("message is %s\n", message)
		if message == "Your_Turn:+" || message == "Your_Turn:-" {
			fmt.Printf("%s\n", message)
			break
		}
	}

	initConn(conn)
	for i := 0; i < 10; i++ {
		conn.Write([]byte("+2a2b\n"))
		message, _ := bufio.NewReader(conn).ReadString('\n')
		log.Print("Message from server: " + message)
	}
}

func taiketsu(conn net.Conn) {

}
