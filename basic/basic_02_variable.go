package basic

import "fmt"

// 变量， 指定存储单元的名称，存储单元存储特定类型的值

func UseCase02Variable() {
	fmt.Println("UseCase02Variable Start")

	// 声明单个变量
	var age int
	// 未给变量赋值，go会自动将其初始化，零值
	fmt.Println("my age is", age)
	// 赋值
	age = 29
	fmt.Println("my new age is", age)

	// 声明变量并初始化
	// var name type = initialValue
	var name string = "victor"
	fmt.Println("my name is", name)

	// 类型推断
	// 如果变量有初始值，那么 Go 能够自动推断具有初始值的变量的类型，可在声明中省略type
	var newAge = 29
	// 自动推它是int类型
	fmt.Println("my new age is", newAge)

	// 声明多个变量
	// var name1, name2 type = initialValue1, initialValue2
	var width, height int = 100, 50
	fmt.Println("width is", width, "height is", height)

	// 在一条声明中声明不通类型的变量
	//var (
	//	name1 = initalValue1
	//	name2 = initialValue2
	//)
	var (
		victor   = "victor"
		city     = "chengdu"
		postcode = 641000
	)
	fmt.Println("my name is", victor, "my city is", city, "my postcode is", postcode)

	// 简短声明 := 操作符
	girlName, girlAge := "mei", 29
	fmt.Println("my girl name is", girlName, " she's age is", girlAge)
	// 要求 := 操作符的左边至少有一个变量是尚未声明的
	// 要求 := 操作符左边的所有变量都有初始值

	fmt.Println("UseCase02Variable End")
	fmt.Println()
}
