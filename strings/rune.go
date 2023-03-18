package strings

import (
	"bytes"
	"fmt"
	"unicode/utf8"
)

func Test() {
	s := "田俊杰"
	fmt.Printf("% x \n", s)
	//内存中的字节序列
	fmt.Println(utf8.RuneCountInString(s))
	for i, r := range "田俊杰" {
		fmt.Printf("%d\t%q\t%d%T\n", i, r, r, r)
	}
	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d\t%c%T\n", i, r, r)
		i += size
	}
	fmt.Printf(string(30000))
	fmt.Printf(string(20426))
	fmt.Printf(string(26480))
	var buf bytes.Buffer
	fmt.Fprintf(&buf,"% x",s)
	buf.WriteRune(30000)
	fmt.Println(buf.String())
}
