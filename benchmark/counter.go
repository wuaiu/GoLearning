package benchmark

import (
	"sync"
	"time"
)

type ICounter interface {
	Count()
	Read()
}
type Counter struct {
	count int
	mutex sync.Mutex
}


func (w *Counter) Count() {
	w.mutex.Lock()
	defer w.mutex.Unlock()
	w.count ++
	time.Sleep(time.Microsecond)
}

func(w *Counter)Read() {
	w.mutex.Lock()
	defer w.mutex.Unlock()
	_ = w.count
	time.Sleep(time.Microsecond)
}
