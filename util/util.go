package util

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
	"time"
)

// ある文字列が出現するまで待ち続ける
func CheckResponse(conn net.Conn, checkString string) string {
	for {
		r := bufio.NewReader(conn)
		for {
			line, err := r.ReadString('\n')
			fmt.Printf("This line:%s\n", line)
			if err == io.EOF {
				break
			} else if err != nil {
				log.Fatal(err)
			}

			if strings.Contains(line, checkString) {
				return line
			}
		}
		time.Sleep(1 * time.Second)
	}
}
