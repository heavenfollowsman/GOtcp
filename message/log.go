package message

import (
	"fmt"
	"github.com/bytedance/sonic"
	"time"
)

/**
日志具体信息
*/

type LogMessage struct {
	// 日志等级
	Level
	// 日志标签（根据这个去区分文件）
	Tag string
	// 控制字符
	ControlCode byte
	// 日志内容
	Content string
	// 时间戳，秒级别
	TimeStamp int64
}

func (m *LogMessage) String() string {
	return fmt.Sprintf("(%s) <%s> [%s]: %s", time.Unix(0, m.TimeStamp).Format("2006-01-02 15:04:05"), m.Tag, m.Level, m.Content)
}

func (m *LogMessage) MarshalTo(to []byte) (int, error) {
	p, err := sonic.Marshal(m)
	if err != nil {
		return 0, err
	}
	l := copy(to, p)
	return l, nil
}

func (m *LogMessage) UnmarshalFrom(buf []byte) error {
	return sonic.Unmarshal(buf, m)
}
