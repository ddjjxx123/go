package main

import (
	"fmt"
	"github.com/ddjjxx123/go/x-server/server"
	"github.com/ddjjxx123/go/x-server/server/impl"
)

type TestRouter struct {
	impl.BaseRouter
}

//前处理
func (router *TestRouter) PreHandle(request server.IXRequest) {
	fmt.Println("Starting TestRouter PreHandle")
	_, err := request.GetConnection().GetConnection().Write([]byte("Pre Handle Success"))
	if err != nil {
		fmt.Println("Pre Handle Err", err)
	}
}

//业务处理
func (router *TestRouter) Handle(request server.IXRequest) {
	fmt.Println("Starting TestRouter Handle")
	_, err := request.GetConnection().GetConnection().Write([]byte(" Handle Success"))
	if err != nil {
		fmt.Println(" Handle Err", err)
	}
}

//后处理
func (router *TestRouter) PostHandle(request server.IXRequest) {
	fmt.Println("Starting TestRouter PostHandle")
	_, err := request.GetConnection().GetConnection().Write([]byte("Pre PostHandle Success"))
	if err != nil {
		fmt.Println("Pre PostHandle Err", err)
	}
}

func main() {
	ixServer := impl.CreateServer()
	ixServer.AddRouter(&TestRouter{})
	ixServer.Serve()

}
