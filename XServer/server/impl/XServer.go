package impl

import (
	"fmt"
	"github.com/ddjjxx123/go/x-server/server"
	"net"
	"time"
)

type XServer struct {
	Name      string
	IPVersion string
	IPAddr    string
	Port      int
	Router    server.IXRouter
}

func (s *XServer) Start() {
	fmt.Printf("Start Serve IP %s:%d \n", s.IPAddr, s.Port)

	//创建连接
	go func() {
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IPAddr, s.Port))
		if err != nil {
			fmt.Println("Resolve TCP Err", err)
			return
		}
		//获取监听
		listener, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("Listen TCP Err", err)
			return
		}

		var connId = CreateConnId()

		for {
			//监听
			tcpConn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("Accept TCP Err", err)
				return
			}

			connection := CreateConnection(tcpConn, connId, s.Router)
			connId++
			//开启连接
			go connection.Start()

		}
	}()

}

func CreateConnId() uint32 {
	//TODO
	return 0
}
func (s *XServer) Stop() {

}
func (s *XServer) Serve() {
	s.Start()
	for {
		time.Sleep(1000)
	}
}

func (s *XServer) AddRouter(router server.IXRouter) {
	s.Router = router
}

func CreateServer(name string) server.IXServer {
	s := &XServer{
		Name:      name,
		IPVersion: "tcp",
		IPAddr:    "127.0.0.1",
		Port:      8888,
	}
	return s
}
