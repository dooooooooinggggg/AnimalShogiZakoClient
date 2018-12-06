package connect

import (
	"fmt"
	"log"
	"net"
	"time"
)

// TCPConnection func
func TCPConnection() net.Conn {
	conn, err := net.Dial("tcp", "shogi.keio.app:80")
	if err != nil {
		log.Fatal(err)
	}
	return conn
}

// KeepAlive func
func KeepAlive(conn net.Conn) {
	fmt.Printf("Keep Alive!\n")
	for {
		conn.Write([]byte(" " + "\n"))
		time.Sleep(30 * time.Second)
	}
}
