package impl

import (
	"alexhades.com/go/server"
	"fmt"
	"net"
	"time"
)

type XServer struct {
	Name      string
	IPVersion string
	IPAddr    string
	Port      int
}

func (s *XServer) Start() {
	fmt.Printf("Start Serve IP%s:%d \n", s.IPAddr, s.Port)

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

		//监听
		tcpConn, err := listener.AcceptTCP()
		if err != nil {
			fmt.Println("Accept TCP Err", err)
			return
		}

		go func() {
			for {
				buf := make([]byte, 512)
				//读
				read, err := tcpConn.Read(buf)
				if err != nil {
					fmt.Println("Read Connection Err", err)
					continue
				}
				//写
				if _, err := tcpConn.Write(buf[:read]); err != nil {
					fmt.Println("Write Err", err)
					continue
				}

			}
		}()
	}()

}
func (s *XServer) Stop() {

}
func (s *XServer) Serve() {
	s.Start()
	for {
		time.Sleep(1000)
	}
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
