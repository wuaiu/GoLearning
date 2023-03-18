package main

import (
	"demo/demo1/strings"
	"fmt"
	"io"
)


func f(out io.Writer) {
	fmt.Printf("%T\n",out)
	out = nil
	fmt.Printf("%T\n",out)
}
func main() {
	//function.Test()
	strings.Test()
	//var w io.Writer
	//fmt.Printf("%T\n",w)
	//w = os.Stdout
	//fmt.Printf("%T\n",w)
	//w = new(bytes.Buffer)
	//fmt.Printf("%T\n",w)
	//w= nil
	//fmt.Printf("%T\n",w)
	//var buf *bytes.Buffer
	//fmt.Printf("%T\n",buf)
	//f(buf)
	//routine.Goroutine()
	//schedule.TestSchedule()
	//hmap.TestMap()
	//var input string
	//fmt.Scanln(&input)
	////fmt.Println("done with:"+ input)
	//str1 := []byte{}
	//var str3 interface{} = str1
	//str2 := str3.(string)
	//str1 := []byte(fmt.Sprint(""))
	//fmt.Println(str1)
	//str2 := string(str1)
	//fmt.Println(str2+"12")
	//if str2== "" {
	//	fmt.Println(666)
	//}
	//map1:= make(map[string]interface{}, 0)
	//value1 := map1["key"]
	//value2 := value1.(bool)
	//value2 ,_:= value1.(bool)
	//if ok {
	//	fmt.Println(value2)
	//}else{
	//	fmt.Println(666)
	//}
	//fmt.Println(value2)
}
