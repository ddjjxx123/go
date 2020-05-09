package impl

import "github.com/ddjjxx123/go/x-server/server"

type XRequest struct {
	Coon server.IXConnection
	Data []byte
}

func (request *XRequest) GetConnection() server.IXConnection {
	return request.Coon
}

func (request *XRequest) GetData() []byte {
	return request.Data
}
