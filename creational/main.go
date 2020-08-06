package main

import "fmt"

func main() {
	animalFactory := AnimalFactory{}
	myAnimal := animalFactory.getAnimal("dog")
	fmt.Println("Animal sound: ", myAnimal.getSound())

	/* So, we need to call our abstract factory to create
	some objects. Go ahead
	*/
	bookFactory := BookFactory{}
	myBook := bookFactory.createBook("spanish", "fantasy")
	fmt.Println(myBook.myNameIs())
	myBook = bookFactory.createBook("portuguese", "fantasy")
	fmt.Println(myBook.myNameIs())
	myBook = bookFactory.createBook("spanish", "drama")
	fmt.Println(myBook.myNameIs())
	myBook = bookFactory.createBook("portuguese", "drama")
	fmt.Println(myBook.myNameIs())
}
