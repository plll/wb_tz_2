package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"
)

// Реализовать постоянную запись данных в канал (главный поток). Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout. Необходима возможность выбора количества воркеров при старте.
//
// Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров.

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()
	ch := make(chan int)
	workersAmount := 100001

	work := func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				return
			case v := <-ch:
				fmt.Println(v)
				return
			}
		}
	}

	go func() {
		for i := 0; i < workersAmount-1; i++ {
			select {
			case <-ctx.Done():
				return
			default:
				go work(ctx)
			}
		}
	}()
	go func() {
		for i := 0; i < workersAmount-1; i++ {
			select {
			case <-ctx.Done():
				return
			default:
				time.Sleep(1000 * time.Millisecond)
				ch <- i
			}

		}
	}()

	<-ctx.Done()
	fmt.Println("Cancelled")

	fmt.Println("Конец программы")
}
