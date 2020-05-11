package server

type IXMessage interface {
	GetMsgId() uint32
	GetDataSize() uint32
	GetData() []byte

	SetMsgId(uint32)
	SetDataSize(uint32)
	SetData([]byte)
}
