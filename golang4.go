package main

import "fmt"

type animalSound interface {
	makeASound() string
}

type Cat struct{}

func (c Cat) makeASound() string {
	return "Meow"
}

type Leon struct{}

func (l Leon) makeASound() string {
	return "Roar"
}

type Tiger struct {
	name string
	age  int
}

func trySound(s animalSound) string {
	return s.makeASound()
}

func main() {

	newCat := Cat{}
	newLeon := Leon{}
	newTiger := Tiger{"yürek", 2}

	fmt.Println(newCat.makeASound())
	fmt.Println(newLeon.makeASound())
	fmt.Println(newTiger)
}
