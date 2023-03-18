package contains

import (
	"fmt"
	"strconv"
	"sync"
)

type ListStruct struct {
	Adj []int64
}
type MapStruct struct {
	m sync.Map
}

func MapTest() {
	//fmt.Println(666)
	//var map1 hmap[string]string = make(hmap[string]string)
	//fmt.Println(len(map1))
	//map1["1"] = ""
	//fmt.Println(map1["1"])
	//_,exist := map1["1"]
	//if exist {
	//	fmt.Println("777")
	//}
	//var list1 []int
	//list1 = append(list1, 3)
	//fmt.Println(list1)
	_, err1 := strconv.ParseInt("w1234", 10, 64)
	if err1 == nil {
		fmt.Println(666)
	}
	str1 := fmt.Sprint(true)
	fmt.Printf("%T\n", str1)
	if str1 == "true" {
		fmt.Println("666")
	}
	requestCounter := 0
	for true {
		fmt.Println(requestCounter)
		if requestCounter > 5000 {
			break
		}
		requestCounter++
	}
	collectionIDMap := make(map[int64][]int64)
	collectionIDMap[7050029497193679621] = []int64{}
	ss := make(map[int64]bool, 0)
	for _, ids := range collectionIDMap {
		for _, id := range ids {
			ss[id] = true
		}
	}
	result := make([]int64, len(ss))
	index := 0
	for key, _ := range ss {
		result[index] = key
		index++
	}
	print(len(result))
	obj1 := &ListStruct{}
	list1 := make([]int64, 0)
	list1 = append(list1, obj1.Adj...)
	fmt.Println()
	fmt.Println(list1)
}
