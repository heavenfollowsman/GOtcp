package message

import (
	"net"
	"testing"
	"time"
)

func TestEncodeDecoder(t *testing.T) {
	t.Skip()
	var edr Encoder
	var conn net.Conn

	m := &LogMessage{
		Content:     "hello world",
		ControlCode: 0b0000,
		Level:       Debug,
		Tag:         "test",
		TimeStamp:   time.Now().UnixNano(),
	}
	b := make([]byte, 1024)
	l, _ := m.MarshalTo(b)

	// 客户端会通过这个接口来向服务端发送Log数据
	_ = edr.Encode(b[:l], conn)
}
