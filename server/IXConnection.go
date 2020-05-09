package server

import "net"

type IXConnection interface {
	//启动连接
	Start()
	//关闭连接
	Stop()
	//获取连接
	GetConnection() *net.TCPConn
	//获取连接ID
	GetConnectionId() uint32
	//获取客户端地址
	GetRemoteAddr() net.Addr
}

//统一业务处理
type HandleFunc func(*net.TCPConn, []byte, int) error
