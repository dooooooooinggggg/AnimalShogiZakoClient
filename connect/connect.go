package connect

import (
	"log"
	"net"
)

// tcp connection
func TcpConnection() net.Conn {
	conn, err := net.Dial("tcp", "shogi.keio.app"+":"+"80")
	if err != nil {
		log.Fatal(err)
	}
	return conn
}

func KeepAlive() {

}
