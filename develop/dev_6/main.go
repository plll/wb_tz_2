package main

import (
	"context"
	"fmt"
	"time"
)

func main1() {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(1*time.Second)) // Создаем контекст, который будет возвращать ctx.Done через какое-то время и отложим 1 секунду
	// Так же можно через любой другой контекст или отловку сигналов системы, но методы все схожие и отлавливают ctx.Done()
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

func main2() {
	quit := make(chan int) // Способ закрытия горутины используя канал извне
	go func() {
		for {
			select {
			case <-quit:
				return
			default:
				//Что-то делаем
			}
		}
	}()
	quit <- 1
}

func main3() {
	go func() {
		for {
			select {
			case <-time.After(2 * time.Second): // Горутина умрет через 2 секунды после начала
				return
			default:
				//Что-то делаем
			}
		}
	}()
}

func main4() {
	ch := make(chan int)
	go func() {
		for {
			val, ok := <-ch
			// В каждом прочтении проверять не закрыл ли кто-нибудь канал и использовать это как знак, что горутине надо умереть
			if !ok {
				return
			}
			fmt.Println(val)
		}
	}()
}
