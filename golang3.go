package main

import "fmt"

type Car struct {
	brandName string
	modelName string
	year      int
}

type Truck struct {
	heigh int
	price int
	car   Car
}

func (truck Truck) newerOrOlder() int {
	if truck.car.year > 2015 {
		return 0
	} else {
		return 1
	}
}

func main() {

	truck := Truck{
		heigh: 3,
		price: 30000,
		car: Car{
			brandName: "Ford",
			modelName: "fMax",
			year:      2019,
		},
	}
	truck2 := Truck{
		heigh: 4,
		price: 20000,
		car: Car{
			brandName: "Ford",
			modelName: "fMax",
			year:      2013,
		},
	}

	fmt.Println(truck)

	anonymousStruct := struct {
		name string
		age  int
	}{
		name: "Burak",
		age:  20,
	}

	if truck.newerOrOlder() == 1 {
		fmt.Println("Old")
	} else {
		fmt.Println("New")
	}
	if truck2.newerOrOlder() == 1 {
		fmt.Println("Old")
	} else {
		fmt.Println("New")
	}

	fmt.Println(anonymousStruct)

}
