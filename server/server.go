package main

import (
	"bufio"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"loggingProject/message"
	"net"
	"os"
)

/**
 	接收到数据 反序列化为结构体实例以后，  先判断控制字段
	控制字段要打印就打印  要保存就保存  如果要忽略 则不管控制字段
	判断完控制字段以后去判断level，根据level的等级决定是打印还是保存
*/

// 服务端结构体 用来保存服务端地址
type Server struct {
	address string
}

// NewServer 创建一个服务端
func NewServer(address string) *Server {
	return &Server{address: address}
}

// Start 开启服务器，并监听制定地址和端口
func (s *Server) Start() {
	listener, err := net.Listen("tcp", s.address)
	if err != nil {
		fmt.Println("服务端开启失败:", err)
		return
	}
	defer listener.Close()

	fmt.Println("服务端开启", s.address)
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("连接失败:", err)
			continue
		}
		// 开启携程
		go s.handleConnection(conn)
	}
}

// handleConnection   先读取要发送的数据的长度，再发送信息 读取信息，打印并保存日志文件
func (s *Server) handleConnection(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)

	for {
		//读取消息长度
		lengthBuf := make([]byte, 4)
		// 这一步lengthBuf中的数值已经改变了 这个函数的返回值是接受的长度
		_, err := io.ReadFull(reader, lengthBuf)

		if err != nil {
			if err == io.EOF {
				// 正常关闭连接时，打印信息并退出循环
				fmt.Println("连接被关闭")
				return
			}
			// 处理其他错误
			fmt.Println("长度读取失败:", err)
			return
		}

		length := binary.BigEndian.Uint32(lengthBuf)
		messageTxt := make([]byte, length)

		_, err = io.ReadFull(reader, messageTxt)
		if err != nil {
			fmt.Println("消息读取失败:", err)
			return
		}

		var log message.LogMessage

		// 网络传输数据的时候 以json格式传输 这里进行反序列化  将json字符串转化为go中的结构体格式
		if err := json.Unmarshal(messageTxt, &log); err != nil {
			fmt.Println("反序列化失败:", err)
			continue
		}
		//
		handleLogMessage(log)
	}
}

// handleLogMessage 如果控制字段不为忽略，则根据控制字段来，忽略则根据level等级来
func handleLogMessage(log message.LogMessage) {
	if log.ControlCode != 0 {
		handleControlCode(log)
	} else {
		handleLevel(log)
	}
}

// handleControlCode 根据控制字符做对应的操作
func handleControlCode(log message.LogMessage) {

	if message.HasOp(log.ControlCode, message.Print) {
		saveLog(log)
	}
	if message.HasOp(log.ControlCode, message.Save) {
		fmt.Println(log.String())
	}
	if message.HasOp(log.ControlCode, message.Ignore) {
		handleLevel(log)
	}
}

// handleLevel 如果需要打印就执行打印，需要保存就执行保存
func handleLevel(log message.LogMessage) {
	fmt.Println(log.Level)
	if log.NeedPrint() {
		fmt.Println(log.String())
	}
	if log.Level.NeedSave() {
		saveLog(log)
	}
}

// saveLog 保存日志文件到文档里。
func saveLog(log message.LogMessage) {
	file, err := os.OpenFile(log.Tag+".log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("保存日志文件失败:", err)
		return
	}

	defer file.Close()

	// go特有的写法 不得不尝一尝
	if _, err := file.WriteString(log.String() + "\n"); err != nil {
		fmt.Println("写入文件失败:", err)
	}
}

func main() {
	address := "localhost:8075"
	server := NewServer(address)
	server.Start()
}
