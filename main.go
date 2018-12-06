package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
	"time"

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

func checkResponse(conn net.Conn, checkString string) string {
	var res string
	for {
		r := bufio.NewReader(conn)
		for {
			line, err := r.ReadString('\n')
			if err == io.EOF {
				break
			} else if err != nil {
				log.Fatal(err)
			}

			if strings.Contains(line, checkString) {
				res = line
				break
			}
		}
		if res != "" {
			break
		}
		time.Sleep(1 * time.Second)
	}
	return res
}

func sendMessage(conn net.Conn) {
	// ここ(対局相手が見つかったタイミング)で，自分が先手か後手か判定したい
	// てかここの初期化の処理書かないと
	// それまでここでストップしないと

	myTurn := checkResponse(conn, "Your_Turn:")
	fmt.Printf("my turn is %s\n", myTurn)
	initConn(conn)

	return

	for i := 0; i < 10; i++ {
		conn.Write([]byte("+2a2b\n"))
		message, _ := bufio.NewReader(conn).ReadString('\n')
		log.Print("Message from server: " + message)
	}
}

func taiketsu(conn net.Conn) {

}
