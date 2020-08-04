package main

import "fmt"

func main() {
	animalFactory := AnimalFactory{}
	myAnimal := animalFactory.getAnimal("dog")
	fmt.Println("Animal sound: ", myAnimal.getSound())
}
