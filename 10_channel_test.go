package experiments

import (
	"testing"
)
//ch := make(chan int, 10) / make(chan int) 赋值 ch <- 1 取值 <- ch
func TestChan(t *testing.T) {
	//1.无缓冲区chan
	// ch := make(chan int)

	// go func() {
	// 	val := <-ch
	// 	fmt.Println(val)
	// }()

	// ch <- 10

	//2.有缓冲区
	ch1 := make(chan int, 10)

	ch1 <- 10
}