package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(1*time.Second)) // Создаем контекст, который будет возвращать ctx.Done через какое-то время и отложим 1 секунду
	ch := make(chan int)
	defer cancel()

	go func() { // Запустим горутину на чтение канала и закроем ее когда будет context.Done()
		for {
			select {
			case <-ctx.Done():
				return
			case v := <-ch:
				fmt.Println(v)
			}
		}
	}()
	for {
		select { //  Добавим селект, для возможности получения ctx.Done() в цикле, чтобы знать когда закрывать работу
		case <-ctx.Done():
			fmt.Println("End")
			close(ch)
			return
		default:
			ch <- 1
		}
	}
}
