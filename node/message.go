package node

import(
	"time"
)

//每条消息的唯一ID
// type MessageID [16]byte
// type MessageID string

//消息结构体
type Message struct {
	ID        string
	Body      []byte
	Timestamp int64
	Attempts  uint16
}

// 创建一条新消息
func NewMessage(id string, body []byte) *Message {
	return &Message{
		ID:        id,
		Body:      body,
		Timestamp: time.Now().UnixNano(),
	}
}