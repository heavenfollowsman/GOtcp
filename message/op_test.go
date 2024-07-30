package message

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOp(t *testing.T) {
	var controlByte byte

	// 啥也没设置  啥权限没有
	assert.False(t, HasOp(controlByte, Print))
	assert.False(t, HasOp(controlByte, Save))
	assert.False(t, HasOp(controlByte, Ignore))
	fmt.Println(controlByte)

	// 设置打印权限，只能打印
	controlByte = SetOp(controlByte, Print)
	assert.True(t, HasOp(controlByte, Print))
	assert.False(t, HasOp(controlByte, Save))
	assert.False(t, HasOp(controlByte, Ignore))
	fmt.Println(controlByte)

	// 设置保存权限，能保存也能打印
	controlByte = SetOp(controlByte, Save)
	assert.True(t, HasOp(controlByte, Print))
	assert.True(t, HasOp(controlByte, Save))
	assert.False(t, HasOp(controlByte, Ignore))
	fmt.Println(controlByte)

	// 设置忽略，
	controlByte = SetOp(controlByte, Ignore)
	assert.True(t, HasOp(controlByte, Print))
	assert.True(t, HasOp(controlByte, Print))
	assert.True(t, HasOp(controlByte, Print))
	fmt.Println(controlByte)
}
