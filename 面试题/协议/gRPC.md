# gRPC

## HTTP 2.0

HTTP 2.0 采用二进制传输，而不是1.1的文本格式。传输快。

1. Stream: 已经建议双向的TCP连接

2. Message：完整的HTTP请求和相应

3. 多路复用。HTTP2.0让所有通信都在一个TCP连接上完成。HTTP 2.0一个连接上可以有任意多个流。