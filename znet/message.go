package znet

type Message struct {
	dataLen uint32
	id      uint32
	data    []byte
}

func (m *Message) GetDataLen() uint32 {
	return m.dataLen
}

func (m *Message) GetMsgId() uint32 {
	return m.id
}

func (m *Message) GetData() []byte {
	return m.data
}

func (m *Message) SetDataLen(len uint32) {
	m.dataLen = len
}

func (m *Message) SetMsgId(id uint32) {
	m.id = id
}

func (m *Message) SetData(bytes []byte) {
	m.data = bytes
}
