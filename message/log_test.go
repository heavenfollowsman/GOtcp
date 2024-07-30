package message

import (
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestLogMessage_Marshal(t *testing.T) {

	m := LogMessage{
		Level:       Error,
		Tag:         "Test",
		ControlCode: 0,
		Content:     "hello world",
		TimeStamp:   time.Now().UnixNano(),
	}

	fmt.Println(m.String())
	b := make([]byte, 1024)
	// 返回的l为序列化后的长度
	l, err := m.MarshalTo(b)
	// 断言 没有错误
	assert.NoError(t, err)
	s, _ := sonic.Marshal(m)
	// l和s的长度相等
	assert.Equal(t, len(s), l)

	m2 := LogMessage{}
	assert.NoError(t, m2.UnmarshalFrom(b[:l]))
	assert.Equal(t, m, m2)
	fmt.Println(m2.String())
}
