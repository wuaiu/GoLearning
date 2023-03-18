package routine

import (
	"fmt"
	"time"
)

func SelectTest() {
	chan1 := make(chan string)
	chan2 := make(chan string)

	go func() {
		time.Sleep(time.Second * 1)
		chan1 <- "one"
	}()
	go func() {
		time.Sleep(time.Second * 2)
		chan2 <- "two"
	}()

	for i:=0; i<2;i++ {
		select {
		case msg1 := <-chan1:
			fmt.Println("received:", msg1)
		case msg2 := <- chan2:
			fmt.Println("received:", msg2)
		}
	}
}
