package server

type IXRouter interface {
	//前处理
	PreHandle(request IXRequest)
	//业务处理
	Handle(request IXRequest)
	//后处理
	PostHandle(request IXRequest)
}
