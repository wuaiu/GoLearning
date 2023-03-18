package benchmark

import (
	"sync"
	"testing"
)

func BenchMarkSliceNoCap(b *testing.B) {
	counter := &Counter{}
	benchmark(b,counter,1000000, 1000)

}

func benchmark(b *testing.B,c ICounter,read,write int) {
	for i :=0 ;i < b.N ;i++ {
		var wg sync.WaitGroup
		for k:=0 ; k < read*16;k++ {
			wg.Add(1)
			go func() {
				c.Read()
				wg.Done()
			}()
		}
		for k:=0 ; k < write*16;k++ {
			wg.Add(1)
			go func() {
				c.Count()
				wg.Done()
			}()
		}
		wg.Wait()
	}

}
