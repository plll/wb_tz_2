package main

import "fmt"

// Разработать программу, которая в рантайме способна определить тип переменной: int, string, bool, channel из переменной типа interface{}.

func printVariableType(x interface{}) {
	_, ok := x.(bool)
	if ok {
		fmt.Println("Boolean")
		return
	}
	_, ok = x.(string)
	if ok {
		fmt.Println("String")
		return
	}
	_, ok = x.(int)
	if ok {
		fmt.Println("Integer")
		return
	}
	_, ok = x.(chan int)
	if ok {
		fmt.Println("Channel")
		return
	}

}

func main() {
	stringVar := "1"
	intVar := 1
	boolVar := true
	var chanVar chan int
	printVariableType(stringVar)
	printVariableType(intVar)
	printVariableType(boolVar)
	printVariableType(chanVar)
}
