package server

type IXServer interface {
	//启动服务
	Start()
	//停止服务
	Stop()
	//开始服务
	Serve()
}
