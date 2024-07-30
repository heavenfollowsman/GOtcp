package message

// uint8和byte好像是一样的

type Op uint8

// 用二进制定义常量 为什么
const (
	// Print 打印这条日志
	Print Op = 0b01
	// Save 保存这条日志
	Save Op = 0b10
	// Ignore 忽略这条日志
	Ignore Op = 0b00
)

// 判断字节是否标志位
func HasOp(b uint8, p Op) bool {
	switch p {
	case Print:
		return b&byte(Print) != 0
	case Save:
		return b&byte(Save) != 0
	case Ignore:
		return b&byte(Ignore) != 0
	default:
		return false
	}
}

func SetOp(b byte, p Op) byte {
	switch p {
	case Print:
		return b | byte(Print)
	case Save:
		return b | byte(Save)
	case Ignore:
		return b | byte(Ignore)
	default:
		return b
	}
}
