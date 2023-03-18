package routine

import "fmt"

func SelectClose() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("reveived job", j)
			} else {
				fmt.Println("received all job ", j)
				done <- true
				return
			}
		}
	}()
	for i := 1; i <= 3; i++ {
		jobs <- i
		fmt.Println("send job ", i)
	}
	close(jobs)
	fmt.Println("send all jobs")
	<-done

}


