package main

import (
	"fmt"
	"time"
)

// Реализовать собственную функцию sleep.

func main() {
	fmt.Println("Start")
	Sleep(10) // Чтобы не использовать функцию time.Sleep мы можем использовать time.After, которая создаст канал, который будет ожидать нужное кол-во секунд до того как в него придет значение разрешающее дальнейшее движение кода
	fmt.Println("End")

}

func Sleep(x int) {
	<-time.After(time.Second * time.Duration(x))
}
