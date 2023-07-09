package znet

import (
	"bytes"
	"encoding/binary"
	"errors"
	"zinx/config"
	"zinx/ziface"
)

type DataPack struct{}

func NewDataPack() *DataPack {
	return &DataPack{}
}

func (d *DataPack) GetHeadLen() uint32 {
	// 自定义消息头长度，id + data len
	return 8
}

func (d *DataPack) Pack(msg ziface.IMessage) ([]byte, error) {
	// 1. 创建缓冲
	dataBuf := bytes.NewBuffer([]byte{})

	// 2. 写数据长度
	if err := binary.Write(dataBuf, binary.LittleEndian, msg.GetDataLen()); err != nil {
		return nil, err
	}

	// 3. 写数据类型id
	if err := binary.Write(dataBuf, binary.LittleEndian, msg.GetMsgId()); err != nil {
		return nil, err
	}
	// 4.写数据
	if err := binary.Write(dataBuf, binary.LittleEndian, msg.GetData()); err != nil {
		return nil, err
	}
	return dataBuf.Bytes(), nil
}

func (d *DataPack) UnPack(binaryData []byte) (ziface.IMessage, error) {
	// 1. 创建一个输入二进制数据的ioReader
	reader := bytes.NewReader(binaryData)

	msg := &Message{}

	// 2. 读取datalen
	if err := binary.Read(reader, binary.LittleEndian, &msg.dataLen); err != nil {
		return nil, err
	}

	// 3. 读取msgID
	if err := binary.Read(reader, binary.LittleEndian, &msg.id); err != nil {
		return nil, err
	}

	// 校样数据长度是否超过运行的最大长度
	if config.Conf.MaxPacketSize > 0 && config.Conf.MaxPacketSize < msg.dataLen {
		return nil, errors.New("too large msg data length")
	}
	return msg, nil
}
