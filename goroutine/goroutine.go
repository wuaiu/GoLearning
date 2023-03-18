package routine

import "fmt"
func f(str string) {
	for i := 0; i<3;i++ {
		fmt.Println(str,":",i)
	}
}

func Goroutine() {
	f("direct")

	go f("goroutine")

	go func(msg string) {
		fmt.Println("msg", msg)
	}("going")

}
