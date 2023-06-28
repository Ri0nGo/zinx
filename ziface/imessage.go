package ziface

// IMessage
// @Description: 消息接口，对通信的消息进行封装
type IMessage interface {
	GetDataLen() uint32
	GetMsgId() uint32
	GetData() []byte

	SetDataLen(uint32)
	SetMsgId(uint32)
	SetData([]byte)
}
