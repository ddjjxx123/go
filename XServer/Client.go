package main

import (
	"fmt"
	"github.com/ddjjxx123/go/x-server/server/impl"
	"io"
	"net"
	"time"
)

func main() {
	go run()
	time.Sleep(5 * time.Second)
}

func run() {
	time.Sleep(3 * time.Second)

	conn, err := net.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println("Connection Err", err)
		return
	}
	for i := 0; i < 10; i++ {

		//发封包message消息
		dataPack := impl.CreateDataPack()
		msg, _ := dataPack.Pack(impl.CreateXMessage(uint32(i), []byte(fmt.Sprint("Client Test Message Time :/d", i))))
		_, err := conn.Write(msg)
		if err != nil {
			fmt.Println("write error err ", err)
			return
		}

		//先读出流中的head部分
		headData := make([]byte, dataPack.GetHeaderLength())
		_, err = io.ReadFull(conn, headData) //ReadFull 会把msg填充满为止
		if err != nil {
			fmt.Println("read head error")
			break
		}
		//将headData字节流 拆包到msg中
		msgHead, err := dataPack.UnPack(headData)
		if err != nil {
			fmt.Println("server unpack err:", err)
			return
		}

		if msgHead.GetDataSize() > 0 {
			//msg 是有data数据的，需要再次读取data数据
			msg := msgHead.(*impl.XMessage)
			msg.Data = make([]byte, msg.GetDataSize())

			//根据dataLen从io中读取字节流
			_, err := io.ReadFull(conn, msg.Data)
			if err != nil {
				fmt.Println("server unpack data err:", err)
				return
			}

			fmt.Println("==> Recv Msg: ID=", msg.MsgId, ", len=", msg.DataSize, ", data=", string(msg.Data))
		}

		time.Sleep(1 * time.Second)
	}
}
