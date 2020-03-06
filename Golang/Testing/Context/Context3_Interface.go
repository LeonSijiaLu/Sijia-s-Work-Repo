package main

import "time"

type Context interface {
	Deadline() (deadline time.Time, ok bool)
	// Deadline获取设置的截止时间，第一个参数返回截止时间，第二个参数ok == false表示有没有设置截止时间
	Done() <-chan struct{}
	// Done返回一个只读的channel, 类型为struct
	// 如果这个channel变得可以读取了，说明parent.Done()已经执行，parent已经发送了取消请求
	Err() error
	// Err返回取消的错误原因
	Value(key interface()) interface{}
}

//
// Context 的继承衍生
//
func WithCancel(parent Context)(ctx Context, cancel CancelFunc)
// 到时间会取消子节点以及下面的所有层
func WithDeadline(parent Context, time.Time) (Context, CancelFunc)
// WithDeadline 和 WithCancel 类似，但是加了时间
func WithTimeout(parent Context, time.Duration)(Context, CancelFunc)
func WithValue(parent Context, key, val interface()) Context
