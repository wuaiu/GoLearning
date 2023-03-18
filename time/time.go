package time

import (
	"fmt"
	"time"
)

func TestTime() {
	ticker := time.NewTicker(time.Millisecond * 500)
	//go func() {
	//	for t := range ticker.C {
	//		if
	//		fmt.Println("Tick at", t.Hour())
	//		fmt.Println("Tick at", t)
	//	}
	//}()
	time.Sleep(time.Millisecond * 1600)
	ticker.Stop()
	fmt.Println("Ticker stopped")
}
