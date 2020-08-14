# Creational Patterns

Currently, this repo contains these patterns:

-   AbstractFactory
-   FactoryMethod
-   Builder

So, if you can run a pattern by once, you can do ... 

```

// factory method
func main() {

	animalFactory := AnimalFactory{}
	myAnimal := animalFactory.getAnimal("dog")
	fmt.Println("Animal sound: ", myAnimal.getSound())
}

```

```

// abstract factory
func main() {

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

```

```

// builder
func main() {

	director := Director{}
	mazdaBuilder := MazdaBuilder{}
	mazdaCar := director.buildMazda(mazdaBuilder)
	fmt.Println(mazdaCar.chasis)

}

```