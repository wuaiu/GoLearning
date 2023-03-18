package main

import "fmt"

func main() {
	arr := []int{1, 2, 3}
	newArr := []*int{}
	for i, _ := range arr {
		newArr = append(newArr, &arr[i])
	}
	for _, v := range arr {
		fmt.Printf("%p\n",&v)
	}
}
