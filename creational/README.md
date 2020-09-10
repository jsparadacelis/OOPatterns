``` 

// builder
func main() {

	director := Director{}
	mazdaBuilder := MazdaBuilder{}
	mazdaCar := director.buildMazda(mazdaBuilder)
	fmt.Println(mazdaCar.chasis)

}

```