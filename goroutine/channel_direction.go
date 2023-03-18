package routine

import "fmt"

func ping(pings chan<- string,msg string) {
	pings <- msg
}

func pong(pings <-chan  string,pongs chan<- string) {
	msg := <- pings
	pongs <- msg
}


func ChannelDirection() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "message")
	pong(pings,pongs)
	msg := <-pongs
	fmt.Println(msg)
}
