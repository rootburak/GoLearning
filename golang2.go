package main

import "fmt"

type Whell struct {
	radius int
}

type Car struct {
	brand      string
	model      string
	year       int
	frontwhell Whell
	backwhell  Whell
}

func main() {
	number := 20
	fmt.Println(manyreturn(number))
	fwhell := Whell{radius: 20}
	bwhell := Whell{radius: 30}
	car1 := Car{
		brand:      "Ford",
		model:      "Focus",
		year:       2014,
		frontwhell: fwhell,
		backwhell:  bwhell,
	}
	fmt.Println(car1)
	firstSlice := []int{1, 2, 3}

	firstSlice = append([]int{0}, firstSlice...)
	fmt.Println(firstSlice)

	firstSlice = append(firstSlice, 4)
	fmt.Println(firstSlice)

	firstSlice = append(firstSlice[:4], append([]int{5, 6, 7}, firstSlice[4:]...)...)
	fmt.Println(firstSlice)
}

func manyreturn(number int) (x, y int) {
	x = number * 10
	y = number * 20
	return
}
