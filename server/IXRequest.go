package server

type IXRequest interface {
	//获取当前连接
	GetConnection() IXConnection
	//获取连接中的数据
	GetData() []byte
}
