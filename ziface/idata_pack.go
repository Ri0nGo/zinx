package ziface

// IDataPack
// @Description: 封装数据包和拆解数据包，解决沾包问题
type IDataPack interface {
	GetHeadLen() uint32
	// 封装数据包，将消息对象封装后发送
	Pack(msg IMessage) ([]byte, error)
	// 将数据流拆解到IMessage对象中
	UnPack([]byte) (IMessage, error)
}
