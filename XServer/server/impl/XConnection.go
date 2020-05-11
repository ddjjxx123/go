package impl

import (
	"errors"
	"fmt"
	"github.com/ddjjxx123/go/x-server/server"
	"io"
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
	IsClosed bool
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

		dataPack := CreateDataPack()

		//读取头信息
		header := make([]byte, dataPack.GetHeaderLength())
		if _, err := io.ReadFull(conn.GetConnection(), header); err != nil {
			fmt.Println("Read Head Err", err)
			break
		}

		//拿到连接，拆包拿到MSG
		message, err := dataPack.UnPack(header)
		if err != nil {
			fmt.Println("UnPack Err", err)
			break
		}
		//放入数据
		var data []byte
		if message.GetDataSize() > 0 {
			data := make([]byte, message.GetDataSize())
			if _, err := io.ReadFull(conn.GetConnection(), data); err != nil {
				fmt.Println("Read Data Err", err)
				break
			}
		}
		//放入数据
		message.SetData(data)
		request := XRequest{
			Coon:    conn,
			Message: message,
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
func (conn *XConnection) SendMessage(msgId uint32, data []byte) error {
	if conn.IsClosed == true {
		return errors.New("Connection closed when send msg")
	}
	//将data封包，并且发送
	dataPack := CreateDataPack()
	msg, err := dataPack.Pack(CreateXMessage(msgId, data))
	if err != nil {
		fmt.Println("Pack error msg id = ", msgId)
		return errors.New("Pack error msg ")
	}

	//写回客户端
	if _, err := conn.Coon.Write(msg); err != nil {
		fmt.Println("Write msg id ", msgId, " error ")
		conn.ExitChan <- true
		return errors.New("conn Write error")
	}

	return nil
}
