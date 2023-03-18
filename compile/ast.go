package main

import (
	"go/ast"
	"go/parser"
	"go/token"
)

func main() {
	src := `
package main
import "fmt"
var a int = 1;
const b int = 2;
func main() {
	for i := 0;i <10; i++{
		c := i + i;
		println("Hello, World!")
	}
}`
	fset := token.NewFileSet()
	f,err := parser.ParseFile(fset,"",src, 0)
	if err != nil {
		panic(err)
	}
	ast.Print(fset, f)

}
