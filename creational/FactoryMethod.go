package main

import "fmt"

type IAnimal interface {
	getSound() string
}

type Dog struct {
}

// implementing methods for Square
func (dog Dog) getSound() string {
	return "Guau"
}

type Cat struct {
}

// implementing method for Circle
func (cat Cat) getSound() string {
	return "Miau"
}

type IAnimalFactory interface {
	getAnimal() IAnimal
}

type AnimalFactory struct {
}

func (animalFactory AnimalFactory) getAnimal(animal string) IAnimal {
	switch animal {
	case "dog":
		fmt.Println("Creating a dog ...")
		return Dog{}
	case "cat":
		fmt.Println("Creating a cat ...")
		return Cat{}
	default:
		return Cat{}
	}
}
