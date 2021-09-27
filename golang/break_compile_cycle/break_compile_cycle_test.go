package break_compile_cycle

import (
	"mygolang/golang/break_compile_cycle/p1"
)

//将代码的依赖转变成成运行时的依赖
func Test_p1p2() {
	pp1 := p1.PP1{}
	pp1.HelloFromP2Side() // Prints: "Hello from package p2"
}
