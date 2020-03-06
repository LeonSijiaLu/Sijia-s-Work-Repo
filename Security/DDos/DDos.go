package main

import (
	"context"
	"flag"
	"net"
	"net/http"
	"time"
)

var (
	method   string
	target   string
	interval int
)

func init() {
	flag.StringVar(&method, "method", "GET", "method used to attack target uri")
	flag.StringVar(&target, "target", "http://news.baidu.com", "target uri for DDoS attack")
	flag.IntVar(&Interval, "interval", 1000, "attacking interval in ms")
}

func httpDial(ctx context.Context, network, addr string) (net.Conn, error) {
	dial := net.Dialer{
		Timeout:   time.Duration(10) * time.Second, //TimeOut是server side, 如果timeout这段时间没有从client来的任何相应，就断开连接
		KeepAlive: time.Duration(60) * time.Second, //Keepalive 是client告诉server, 不要再第一个request之后就断开连接，这个连接应该有keepalive的长度
	}
	conn, err := dial.Dial(network, addr)
	return conn, err
}

func newHttpClient() *http.Client {

}
