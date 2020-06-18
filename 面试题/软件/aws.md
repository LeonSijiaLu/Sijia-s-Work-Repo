# AWS

## ALB 和 NLB

AWS 有三种负载均衡。Application, Network, Classic

1. ALB 支持 HTTP 负载均衡，NLB 支持 TCP 负载均衡。

2. ALB 支持 HTTP2 和 WebSocket。并且支持动态端口分配，适用于容器

3. NLB 可以发送到制定实例的端口