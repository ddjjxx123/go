package impl

import (
	"fmt"
	"github.com/ddjjxx123/go/server"
	"net"
)

type XConnection struct {
	//连接套接字
	Coon *net.TCPConn
	//sessionId
	CoonId int32
	//关闭标识
	IsClose bool
	//业务处理api
	HandApi server.HandleFunc
	//传输标识通道
	ExitChan chan bool
}

func (conn *XConnection) StartingReader() {
	fmt.Println("Starting Reader")
	defer fmt.Println(conn.GetRemoteAddr().String(), "not Exist")
	defer conn.Stop()
	//读取数据
	for {
		buf := make([]byte, 512)
		read, err := conn.GetConnection().Read(buf)
		if err != nil {
			fmt.Println("Read Connection Err", err)
			//读取错误告诉连接关闭
			conn.ExitChan <- true
			continue
		}
		//读取正确执行业务api
		if conn.HandApi(conn.Coon, buf, read) != nil {
			fmt.Println("Hand Api Err", err)
			//关闭连接
			conn.ExitChan <- true
			return
		}
	}
}

func (conn *XConnection) Start() {
	//开启读取数据
	go conn.StartingReader()
	for {
		select {
		//如果监听到退出请求，不再阻塞
		case <-conn.ExitChan:
			return
		}
	}

}
func (conn *XConnection) Stop() {
	if conn.IsClose {
		//如果已经关闭 返回
		return
	}
	conn.IsClose = false
	//关闭连接
	conn.Coon.Close()
	//告诉需要处理的channel
	conn.ExitChan <- true
	//关闭Chan
	close(conn.ExitChan)

}
func (conn *XConnection) GetConnection() *net.TCPConn {
	return conn.Coon
}
func (conn *XConnection) GetConnectionId() int32 {
	return conn.CoonId

}
func (conn *XConnection) GetRemoteAddr() net.Addr {
	return conn.Coon.RemoteAddr()
}

func CreateConnection(conn *net.TCPConn, coonId int32, handleApi server.HandleFunc) *XConnection {
	fmt.Printf("Create Connection RemoteAddr:%s CoonId:%d\n", conn.RemoteAddr().String(), coonId)
	return &XConnection{Coon: conn, CoonId: coonId, IsClose: false, HandApi: handleApi, ExitChan: make(chan bool, 1)}
}
