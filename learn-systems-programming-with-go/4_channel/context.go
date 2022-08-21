package main

import (
	"context"
	"fmt"
)

func main() {
	fmt.Println("start sub()")
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		fmt.Println("sub() is finished")
		cancel()
	}()

	<-ctx.Done()
	fmt.Println("all tasks are finished")
}

// チャネルだけでキャンセル処理する場合
func channelImple() {
	fmt.Println("start sub()")
	done := make(chan struct{})

	go func() {
		fmt.Println("sub() is finished")
		close(done)
	}()
<-done
	fmt.Println("all tasks are finished")
}
