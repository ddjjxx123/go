package main

import (
	"fmt"
	"github.com/ddjjxx123/go"
	"github.com/ddjjxx123/go/blob/master/server/impl"
	"net"
	"time"
)

func main() {
	go test()
	impl.CreateServer("XServer1").Serve()
}

func test() {
	time.Sleep(3 * time.Second)
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("Connection Err", err)
		return
	}
	for i := 0; i < 100; i++ {
		conn.Write([]byte(fmt.Sprintf("Test Test %d", i)))
		buf := make([]byte, 512)
		read, _ := conn.Read(buf)
		fmt.Printf("Response=%s read=%d \n", buf, read)
		time.Sleep(100 * time.Millisecond)
	}
}
