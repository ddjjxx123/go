package main

import (
	"fmt"
	"net"
	"strings"
	"time"
)

func main() {
	go test()
	time.Sleep(5 * time.Second)
}

func test() {
	time.Sleep(3 * time.Second)
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("Connection Err", err)
		return
	}
	for i := 0; i < 10; i++ {

		conn.Write([]byte(strings.TrimSpace(fmt.Sprintf("Test Test %d", i))))
		buf := make([]byte, 512)
		read, _ := conn.Read(buf)

		fmt.Printf("Response=%s read=%d\n", buf[:read], read)
		time.Sleep(100 * time.Millisecond)
	}
}
