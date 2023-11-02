package main

import (
	"fmt"
	"sync"
	"time"
)

//Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде. По завершению программа должна выводить итоговое значение счетчика.

type Counter struct {
	inc int
	mu  sync.Mutex
}

func main() {
	c := Counter{inc: 0}
	go addTen(&c)
	go addTen(&c)
	time.Sleep(time.Second * 1)
	fmt.Println(c.inc)
}

func addTen(c *Counter) {
	for i := 0; i < 10; i++ {
		c.mu.Lock()
		c.inc += 1
		c.mu.Unlock()
	}
}
