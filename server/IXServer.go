package server

type IXServer interface {
	Start()
	Stop()
	Serve()
}
