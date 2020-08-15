package main

import "fmt"

func isEven(i int) bool {
	return i%2 == 0
}

func getPrize() (int, string) {
	i := 2
	s := "goldfish"
	return i, s
}

func sumNumbers(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func sayHi() (x, y string) {
	x = "hello"
	y = "world"
	return
}

func feedMe(portion int, eaten int) int {
	eaten = portion + eaten
	if eaten >= 5 {
		fmt.Printf("I'm full! I've eaten %d\n", eaten)
		return eaten
	}
	fmt.Printf("I'm still hungry! I've eaten %d\n", eaten)
	return feedMe(portion, eaten)
}

func anotherFunction(f func() string) string {
	return f()
}

// func main() {
// 	fmt.Println("4.1:调用函数")
// 	fmt.Printf("%v\n", isEven(123))
// 	fmt.Printf("%v\n", isEven(322))
// 	fmt.Println("4.2:函数返回多个值")
// 	quantity, prize := getPrize()
// 	fmt.Printf("You won %v %v\n", quantity, prize)
// 	fmt.Println("4.3:使用不定参数函数")
// 	result := sumNumbers(1, 2, 3, 4)
// 	fmt.Printf("The result is %v\n", result)
// 	fmt.Println("4.4:具名返回值")
// 	fmt.Println(sayHi())
// 	fmt.Println("4.5:递归函数")
// 	feedMe(1, 0)
// 	fmt.Println("4.6:将函数作为参数传递")
// 	fn := func() {
// 		fmt.Println("function called")
// 	}
// 	fn()
// 	fmt.Println("4.7:函数作为值")
// 	fn1 := func() string {
// 		return "function called"
// 	}
// 	fmt.Println(anotherFunction(fn1))
// }
