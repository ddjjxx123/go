package impl

import "github.com/ddjjxx123/go/x-server/server"

type BaseRouter struct {
}

//前处理
func (router *BaseRouter) PreHandle(request server.IXRequest) {}

//业务处理
func (router *BaseRouter) Handle(request server.IXRequest) {}

//后处理
func (router *BaseRouter) PostHandle(request server.IXRequest) {}
