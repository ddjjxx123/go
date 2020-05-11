package impl

import (
	"bytes"
	"encoding/binary"
	"errors"
	"github.com/ddjjxx123/go/x-server/config"
	"github.com/ddjjxx123/go/x-server/server"
)

type DataPack struct {
}

func CreateDataPack() *DataPack {
	return &DataPack{}
}

/*
	获取请求头数据长度
*/
func (dataPack *DataPack) GetHeaderLength() uint32 {
	//Id uint32(4字节) +  Size uint32(4字节)
	return 8
}

/*
	封包
*/
func (dataPack *DataPack) Pack(message server.IXMessage) ([]byte, error) {
	write := bytes.NewBuffer([]byte{})
	//写入id
	if err := binary.Write(write, binary.LittleEndian, message.GetMsgId()); err != nil {
		return nil, err
	}
	//写入dataSize
	if err := binary.Write(write, binary.LittleEndian, message.GetDataSize()); err != nil {
		return nil, err
	}
	//写入data
	if err := binary.Write(write, binary.LittleEndian, message.GetData()); err != nil {
		return nil, err
	}

	return write.Bytes(), nil
}

/*
	拆包
*/
func (dataPack *DataPack) UnPack(data []byte) (server.IXMessage, error) {
	message := &XMessage{}
	reader := bytes.NewReader(data)
	//读msgID
	if err := binary.Read(reader, binary.LittleEndian, &message.MsgId); err != nil {
		return nil, err
	}
	//读数据长度
	if err := binary.Read(reader, binary.LittleEndian, &message.DataSize); err != nil {
		return nil, err
	}
	//判断是否超过最大长度
	if config.GlobalServerObject.MaxDataSize > 0 && message.DataSize > config.GlobalServerObject.MaxDataSize {
		return nil, errors.New("Too Large MSG Data Recieved")
	}

	return message, nil
}
