package main

import (
	"time"
	"fmt"
	"sync"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func sum(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	c <- sum // 将和送入c
}

//func fibonacci(n int, c chan int) {
//	x, y := 0, 1
//	for i := 0; i < n; i++ {
//		c <- x
//		x, y = y, x + y
//	}
//	close(c)
//}

func fibonacci(c, quit chan  int) {
	x, y := 0, 1
	for {
		select {
		case c <- x :
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}
// SafeCounter 的并发使用是安全的
type SafeCounter struct {
	v map[string]int
	mux sync.Mutex
}

// Inc 增加给定key的计数器的值
func (c *SafeCounter) Inc(key string) {
	c.mux.Lock()
	// Lock 之后同一时刻只能有一个goroutine 能访问c.v
	c.v[key]++
	c.mux.Unlock()
}

func (c *SafeCounter) Value(key string) int {
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.v[key]
}



func main() {
	//go say("world")
	//say("hello")

	//a := []int{7, 2, 8, -9, 4, 0}
	//c := make(chan  int)
	//go sum(a[: len(a)/2], c)
	//go sum(a[len(a)/2:], c)
	//x ,y := <-c, <-c // 从c中获取
	//
	//fmt.Println(x, y, x+y)

	//ch := make(chan  int, 2)
	//ch <- 1
	//ch <- 2
	//fmt.Println(<-ch)
	//fmt.Println(<-ch)

	//c := make(chan int, 10)
	//go fibonacci(cap(c), c)
	//for i := range  c { // 不断从channel接收值，直到它被关闭
	//	fmt.Println(i)
	//}

	//c := make(chan  int)
	//quit := make(chan int)
	//go func() {
	//	for  i := 0; i < 10; i++ {
	//		fmt.Println(<-c)
	//	}
	//	quit <- 0
	//}()
	//fibonacci(c, quit)

	//tick := time.Tick(100 * time.Millisecond)
	//boom := time.After(500 * time.Millisecond)
	//for {
	//	select {
	//	case <-tick:
	//		fmt.Println("tick.")
	//	case <-boom:
	//		fmt.Println("Boom!")
     //       return
	//	default:
	//		fmt.Println("    .")
	//		time.Sleep(50 * time.Millisecond)
	//	}
	//}

	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("someKey")
	}
	time.Sleep(time.Second)
	fmt.Println(c.Value("someKey"))
}
