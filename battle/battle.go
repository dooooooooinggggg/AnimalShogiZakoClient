package battle

import "net"

var myTurn int

// ここではライフサイクルの管理
func Battle(conn net.Conn, myTurnString string) {
	switch myTurnString {
	case "Your_Turn:+":
		myTurn = 1
	case "Your_Turn:-":
		myTurn = 0
	}
}
