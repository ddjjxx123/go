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
	CoonId uint32
	//关闭标识
	IsClose bool
	//路由
	Router server.IXRouter
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
		_, err := conn.GetConnection().Read(buf)
		if err != nil {
			fmt.Println("Read Connection Err", err)
			//读取错误告诉连接关闭
			conn.ExitChan <- true
			continue
		}
		request := XRequest{
			Coon: conn,
			Data: buf,
		}
		//执行业务方法
		go func(xRequest server.IXRequest) {
			conn.Router.PreHandle(xRequest)
			conn.Router.Handle(xRequest)
			conn.Router.PostHandle(xRequest)
		}(&request)

	}
}

func CreateConnection(conn *net.TCPConn, coonId uint32, router server.IXRouter) *XConnection {
	fmt.Printf("Create Connection RemoteAddr:%s CoonId:%d\n", conn.RemoteAddr().String(), coonId)
	return &XConnection{Coon: conn, CoonId: coonId, IsClose: false, Router: router, ExitChan: make(chan bool, 1)}
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
func (conn *XConnection) GetConnectionId() uint32 {
	return conn.CoonId

}
func (conn *XConnection) GetRemoteAddr() net.Addr {
	return conn.Coon.RemoteAddr()
}
