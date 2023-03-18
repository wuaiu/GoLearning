package function

import "fmt"

type Animal struct {
}

func (animal *Animal) Run() {
	fmt.Println("run ...")
}

func (animal *Animal) String() string {
	//fmt.Println("run ...")
	return "I'm a animal"
}

func Put(animal *Animal) {
	fmt.Printf("%T\n", animal)
	animal.Run()
}

func Test() {
	animal := Animal{}
	Put(animal)
	fmt.Println(&animal)
}
