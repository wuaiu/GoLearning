package sync

import (
	"fmt"
	"sync"
)

func SyncTest() {
	var m sync.Map
	m.Store("test", 18)
	m.Store("mo", 20)
	age, _ := m.Load("test")
	fmt.Println(age.(int))
	m.Range(func(key, value interface{}) bool {
		name := key.(string)
		age := value.(int)
		fmt.Println(name, age)
		return true

	})
}
