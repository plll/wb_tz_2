package main

import (
	"fmt"
	"sync"
	"time"
)

// Реализовать конкурентную запись данных в map.

type ConcurrentMap struct {
	m  map[string]int
	mu sync.Mutex
}

func main() {
	c := ConcurrentMap{m: make(map[string]int)}
	c.m["test"] = 0
	for i := 0; i < 50; i++ { // Создаем 50 горутин, которые каждая добавят по 10 по 1 за раз
		go addTenToMapField(&c)
	}
	time.Sleep(time.Second * 1)
	fmt.Println(c.m["test"])
}

func addTenToMapField(c *ConcurrentMap) {
	for i := 0; i < 10; i++ {
		c.mu.Lock() // Блокируем остальным горутинам доступ
		c.m["test"] += 1
		c.mu.Unlock() // Открываем доступ для других горутин
	}
}
