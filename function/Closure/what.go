package main

import "fmt"

/* 其中defer语句延迟执行了一个匿名函数，因为这个匿名函数捕获了外部函数的局部变量v，
这种函数我们一般叫闭包。闭包对捕获的外部变量并不是传值方式访问，而是以引用的方式访问。 */
func Inc() (v int) {
	defer func() { v++ }()
	return 42
}

func main() {
	fmt.Println(Inc())

	//加深理解
	for i := 0; i < 3; i++ {
		defer func() { println(i) }()
	}

	for i := 0; i < 3; i++ {
		// 通过函数传入i
		// defer 语句会马上对调用参数求值
		defer func(i int) { println(i) }(i)
	}
}
