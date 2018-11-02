package basic

import "fmt"

// 常量

func UseCase04Constant() {
	fmt.Println("UseCase04Constant Start")

	// 常量表示固定值， 不能重新赋值
	// const 关键字
	// 常量的值会在编译的时候确定
	{
		const a = 55
		fmt.Println("a value is", a)
		//a = 33 // cannot assign to constant a
	}

	// 字符串常量
	// 双引号中的任何值都是 Go 中的字符串常量
	{
		const hello = "hello world"
		fmt.Println("hello constant", hello)
	}

	// 布尔常量
	// 只有2个值true和false

	// 数字常量
	// 整数、浮点、复数

	// 常量

	fmt.Println("UseCase04Constant End")
	fmt.Println()
}
