package impl

type XMessage struct {
	MsgId    uint32 //数据id
	DataSize uint32 //数据长度
	Data     []byte //值
}

func CreateXMessage(msgId uint32, data []byte) *XMessage {
	return &XMessage{
		MsgId:    msgId,
		DataSize: uint32(len(data)),
		Data:     data,
	}
}

func (msg *XMessage) GetMsgId() uint32 {
	return msg.MsgId
}
func (msg *XMessage) GetDataSize() uint32 {
	return msg.DataSize
}
func (msg *XMessage) GetData() []byte {
	return msg.Data
}

func (msg *XMessage) SetMsgId(msgId uint32) {
	msg.MsgId = msgId
}
func (msg *XMessage) SetDataSize(dataSize uint32) {
	msg.DataSize = dataSize
}
func (msg *XMessage) SetData(data []byte) {
	msg.Data = data
}
