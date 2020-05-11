package main

import (
	"github.com/streadway/amqp"
	"bytes"
	"fmt"
)

// https://www.jianshu.com/p/434eb8bfaa5f

// Header
// Topic: 根据 queue 的某种 pattern, 比如有的 queue.US.1, 有的 queue.EU.1
// Fan Out: 所有的 client 都会通过 queue 收到信息
// Direct: 直接发给 queue1

// RabbitMQ uses amqp protocol

const (
	queueName = "push.msg.q"
	exchange = "t.msg.ex"
	mqurl = "amqp://guest:guest@192.168.157.20:5672"
)

func connectRabbit() *amqp.Channel{
	conn, _ := amqp.Dial(mqurl)
	ch, _ := conn.Channel()
	return ch
}

func Close(conn *amqp.Connection){
	conn.Close()
}

func push(ch *amqp.Channel, exchangeName string, qName string, exchangeType string, durable bool, autoDelete bool, internal bool) error{
	if ch == nil {ch = connectRabbit()}
	msg := "Test Message"
	err := ch.ExchangeDeclare(exchangeName, exchangeType, durable, autoDelete, internal, false, nil)
	if err != nil {return err}

	_, err = ch.QueueDeclare(qName, false, false, false, false, nil)
	if err != nil {return err}

	err = ch.QueueBind(qName, "info", exchangeName, false, nil)
	if err != nil {return err}

	err = ch.Publish(exchangeName, "info", false, false, amqp.Publishing{
		ContentType:"text/plain", Body:[]byte(msg),
	})
	if err != nil {return err}
	return nil
}

func BytesToString(b *[]byte) string{
	s := bytes.NewBuffer(*b)
	return s.String()
}

func receive(ch *amqp.Channel, qName string) (error, string){
	msg, _, err := ch.Get(qName, false)
	if err != nil {return err, ""}
	err = ch.Ack(msg.DeliveryTag, false)
	if err != nil {return err, ""}
	return nil, BytesToString(&(msg.Body))
}

func main(){
	ch := connectRabbit()
	err := push(ch, "exchange_A", "q1", "direct", false, false, false)
	if err != nil {panic(err)}
	err, content := receive(ch, "q1")
	if err != nil {panic(err)}
	fmt.Println("content is ", content)
}
