package main

import "fmt"

var s9 = "3.9 常用变量"
var s10 = "3.10 变量作用域"

const greeting string = "3.14 常量"

func showMemoryAddress12(x int) {
	fmt.Println(&x)
	return
}
func showMemoryAddress13(x *int) {
	fmt.Println(x)
	return
}

func main() {
	// 3.1 声明变量
	var s1 string = "3.1 test"
	fmt.Println(s1)
	// 3.2 声明后再赋值
	var s2 string
	s2 = "3.2 test"
	fmt.Println(s2)
	// 3.3
	var i3 int
	i3 = 1
	fmt.Println(i3)
	// 3.4
	var s4, t4 string = "foo", "bar"
	fmt.Println(s4)
	fmt.Println(t4)
	// 3.5
	var (
		s5 string = "foo"
		t5 int    = 4
	)
	fmt.Println(s5)
	fmt.Println(t5)
	// 3.6
	var i6 int
	var f6 float64
	var b6 bool
	var s6 string
	fmt.Printf("%v %v %v %q\n", i6, f6, b6, s6)
	// 3.7
	var s7 string
	if s7 == "" {
		fmt.Println("s has not been assigned a value and is zero valued")
	}
	// 3.8 简短变量
	s8 := "3.8 简短变量"
	fmt.Println(s8)
	// 3.9
	i9 := 42
	fmt.Println(s9)
	fmt.Println(i9)
	// 3.10
	fmt.Printf("Printing `s10` variable from outer block %v\n", s10)
	b10 := true
	if b10 {
		fmt.Printf("Printing `b10` variable from outer block %v\n", b10)
		i10 := 42
		if b10 != false {
			fmt.Printf("Printing `i10` variable from outer block %v\n", i10)
		}
	}
	// 3.11 打印内存地址
	s11 := "Hello world"
	fmt.Println(&s11)
	// 3.12 将变量作为值传递
	i12 := 1
	fmt.Println(&i12)
	showMemoryAddress12(i12)
	// 3.13 将变量作为指针传递
	i13 := 1
	fmt.Println(&i13)
	showMemoryAddress13(&i13)
	// 3.14  常量
	fmt.Println(greeting)
}
