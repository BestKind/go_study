## 问题
1. 总结几种 socket 粘包的解包方式: fix length/delimiter based/length field based frame decoder。尝试举例其应用
2. 实现一个从 socket connection 中解码出 goim 协议的解码器。

## 回答
### 粘包
粘包原因：
1. 发送数据时，数据长度超过缓冲区间，同一个数据包就要通过多次发送完成，最终形成半包，
当数据长度小于缓冲区间，同一个数据包填满后包含了多个数据包内容，最终形成了粘包
2. 接受数据时，读取缓冲区的数据流不及时，导致缓冲区放了多个数据包内容，再次读取时形成了粘包

解决方式：
1. fix length
每次发送固定长度数据包，且不超过缓冲区，接收方每次按照固定长度接收数据

2. delimiter based
发送数据包结尾处添加特殊字符，用来标记数据边界

3. length field based frame decoder
在数据包包头添加包长度信息

### goim
协议结构
4bytes PacketLen 包长度
2bytes HeaderLen 头长度
2bytes Version 协议版本号
4bytes Operation 协议指令
4bytes Sequence 序列号
PacketLen-HeaderLen Body 实际业务数据
