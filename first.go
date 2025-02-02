package main

import (
	"fmt"
	"strings"
	"time"
)

type person struct {
	name string
	age  int
}

func sayhello(name string) string {
	if strings.Contains(name, "burak") {
		fmt.Println("Hello BURAK")
	}
	strings.ToLower(name)
	fmt.Println(strings.ToUpper(name))
	return strings.ToUpper(name)
}

func main() {
	fmt.Println("Hello Go")
	name := "burak"
	fmt.Println(name)
	var x int
	fmt.Scanf("%d", &x)
	if 11 != 10 {
		println("not equals")
	}
	switch x {
	case 1:
		fmt.Println("x equals 1 ")
	case 2:
		fmt.Println("x equals 2")
	case 3:
		fmt.Println("x equals 3")
	default:
		fmt.Println("unkown number ")
	}
	for i := 0; i <= x; i++ {
		fmt.Println("current number %d ", i)
	}
	p1 := person{name: "burak", age: 20}
	fmt.Println(p1)
	fmt.Println(sayhello(p1.name))

	go asyncc()
	for i := 0; i < 5; i++ {

		fmt.Println("Main Async Hello!")
		time.Sleep(500 * time.Millisecond)

	}

	if length := getLength(p1.name); length > 4 {
		fmt.Println("bigger than 4")
	}

}

func asyncc() {
	for i := 0; i < 5; i++ {
		fmt.Println("Func Async Hello!")
		time.Sleep(500 * time.Millisecond)
	}
}

func getLength(name string) int {
	return len(name)
}
