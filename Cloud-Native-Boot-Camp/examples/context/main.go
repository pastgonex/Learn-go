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

func main() {
	baseCtx := context.Background()
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

}
