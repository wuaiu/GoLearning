package schedule

import (
	"fmt"
	"runtime"
	"time"
)

func TestSchedule() {
	var x int
	threads := runtime.GOMAXPROCS(0)

	for i:=0 ; i< threads;i++ {
		go func() {
			for {
				x++
			}
		}()
	}
	time.Sleep(time.Second)
	fmt.Println("x=",x)
}
