package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"loggingProject/message"
	"net"
	"time"
)

type MyEncoder struct{}

// 实现Encode接口 输入信息和连接  返回错误信息
func (s *MyEncoder) Encode(data []byte, w io.Writer) error {

	// 发送长度数据
	length := uint32(len(data))
	lengthBuf := make([]byte, 4)

	// 好像这样能防止粘包，虽然不知道原理
	binary.BigEndian.PutUint32(lengthBuf, length)

	// 发送长度数据
	_, err := w.Write(lengthBuf)
	if err != nil {
		return err
	}

	// 发送具体数据
	_, err = w.Write(data)
	return err
}

func main() {

	conn, err := net.Dial("tcp", "localhost:8075")
	if err != nil {
		fmt.Println("连接服务器失败:", err)
		return
	}
	defer conn.Close()

	// 创建一个 LogMessage 实例
	m := &message.LogMessage{
		Content:     "hello world",
		ControlCode: 0b01,
		Level:       message.Warn,
		Tag:         "test",
		TimeStamp:   time.Now().UnixNano(),
	}

	// 序列化 LogMessage
	b := make([]byte, 1024)
	l, err := m.MarshalTo(b)
	if err != nil {
		fmt.Println("序列化失败:", err)
		return
	}

	// 创建 MyEncoder 实例
	encoder := MyEncoder{}

	// 使用 MyEncoder 将数据发送到服务器
	err = encoder.Encode(b[:l], conn)
	if err != nil {
		fmt.Println("发送数据失败:", err)
	}

}
