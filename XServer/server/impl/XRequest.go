package impl

import "github.com/ddjjxx123/go/x-server/server"

type XRequest struct {
	Coon    server.IXConnection
	Message server.IXMessage
}

func (request *XRequest) GetConnection() server.IXConnection {
	return request.Coon
}

func (request *XRequest) GetData() []byte {
	return request.Message.GetData()
}
