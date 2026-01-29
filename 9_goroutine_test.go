package experiments

import (
	"fmt"
	"testing"
	"time"
)

//A goroutine is a lightweight thread of execution.
//协程是一种轻量级的执行线程
func printlnHelloWorld() {
	fmt.Println("Hello,World!")
}

func TestGoroutine(t *testing.T) {

	go printlnHelloWorld()

	go func() {
		fmt.Println("Hello again!")
	}()

	time.Sleep(time.Second)
}