package server

type IXDataPack interface {
	GetHeaderLength() uint32                //获取请求头数据长度
	Pack(message IXMessage) ([]byte, error) //封包
	UnPack(data []byte) (IXMessage, error)  //拆包
}
