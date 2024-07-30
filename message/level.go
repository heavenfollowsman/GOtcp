package message

// Level 日志等级
type Level int

// 日志等级 1234
const (
	Debug Level = iota + 1
	Info
	Warn
	Error
)

// NeedPrint 是否可以打印
func (l Level) NeedPrint() bool {
	return l >= Info
}

// NeedSave 是否可以保存
func (l Level) NeedSave() bool {
	return l >= Warn
}

func (l Level) String() (str string) {
	switch l {
	case Debug:
		str = "Debug"
	case Info:
		str = "Info"
	case Warn:
		str = "Warn"
	case Error:
		str = "Error"
	default:
		str = "Unknown"
	}
	return
}
