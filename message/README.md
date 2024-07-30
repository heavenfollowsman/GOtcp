
## Log服务

### 需求

- 客户端连接发来LogMessage结构体
- 通过LogMessage的Level字段判断是否打印或者保存日志，日志保存的文件名为LogMessage的Tag字段
- LogMessage中的ControlCode字段为控制字符，其作用高于一切（控制字段表明打印就打印，保存就保存，忽略就忽略，无视Level）
- 提供一个protocol.go文件下接口的实现（服务端调用接口的方式会在protocol_test.go内有示例）