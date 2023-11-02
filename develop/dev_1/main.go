package main

//Дана структура Human (с произвольным набором полей и методов). Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).

type Delivery struct {
}

type Order struct {
	Delivery // Процесс встраивания (наследование)
}

func main() {
}
