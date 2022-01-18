package main

import (
	"context"
	"fmt"
	"time"
)

//type Context interface {
//	Deadline() (deadline time.Time, ok bool)
//	Done() <-chan struct{}
//	Err() error
//	Value(key interface{}) interface{}
//}

/*
超时、取消操作或者一些异常情况， 往往需要进行抢占操作或者中断后续操作
Context 是设置截止日期、同步信号，传递请求相关值的结构体
context.Background()
	Background通常被用于主函数、初始化以及测试中， 作为一个顶层的context， 也就是说一般我们创建的context都是基于Background
context.TODO()
	TODO是在不确定使用什么context的时候才会使用
context.WithDeadline()
	超时时间
context.WithValue()
	向context添加键值对
context.WithCancel()
	创建一个可取消的context
*/

func main() {
	baseCtx := context.Background() // 通常被用于主函数、初始化以及测试中， 作为一个顶层的context， 也就是说一般我们创建的context都是基于Backgroud
	ctx := context.WithValue(baseCtx, "a", "b")
	go func(c context.Context) {
		fmt.Println(c.Value("a"))
	}(ctx)
	timeoutCtx, cancel := context.WithTimeout(baseCtx, time.Second)
	defer cancel()
	go func(ctx context.Context) {
		ticker := time.NewTicker(1 * time.Second)
		for _ = range ticker.C {
			select {
			case <-ctx.Done():
				fmt.Println("child process exited")
				return
			default:
				fmt.Println("enter  default")
			}
		}

	}(timeoutCtx)
	select {
	case <-timeoutCtx.Done():
		time.Sleep(1 * time.Second)
		fmt.Println("main process exit!")
	}
	// time.Sleep(5 * time.Second)
}
