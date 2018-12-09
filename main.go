package main

import (
	"bufio"
	"fmt"
	"net"

	"github.com/dooooooooinggggg/AnimalShogiZakoClient/battle"
	"github.com/dooooooooinggggg/AnimalShogiZakoClient/connect"
	"github.com/dooooooooinggggg/AnimalShogiZakoClient/util"
)

func main() {
	conn := connect.TCPConnection()
	defer conn.Close()
	fmt.Printf("Connected to 'Animal Shogi Server'\n")
	go connect.KeepAlive(conn)

	myTurn := util.CheckResponse(conn, "Your_Turn:")
	// fmt.Printf("my turn is %s\n", myTurn)
	util.CheckResponse(conn, "END Game_Summary")
	initConn(conn)
	battle.Battle(conn, myTurn)
}

func initConn(conn net.Conn) {
	conn.Write([]byte("LOGIN:Saikyou_AI\n"))
	conn.Write([]byte("AGREE\n"))
	message, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Printf("%s\n", message)
}
