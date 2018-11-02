package basic

import "fmt"

// 指针
// 指针是一种存储变量内存地址（Memory Address）的变量

// 指针的声明
// 指针变量的类型为 *T，该指针指向一个 T 类型的变量。


func change(val *int) {
	*val = 55
}

func UseCase11Pointer()  {
	fmt.Println("UseCase11Pointer Start")

	b := 255
	var a *int = &b
	//  *int
	fmt.Printf("Type of a is %T\n", a)
	//  0xc0000583f8
	fmt.Println("address of b is", a)

	// 指针的零值 nil
	{
		a := 25
		// b is nil
		var b *int
		if b == nil {
			// b is <nil>
			fmt.Println("b is", b)
			b = &a
			// b initialization  is 0xc000058400
			fmt.Println("b initialization  is", b)
		}
	}

	// 指针的解引用
	// 指针的解引用可以获取指针所指向的变量的值。将 a 解引用的语法是 *a
	{
		b := 255
		a := &b
		// address of b is 0xc000058408
		fmt.Println("address of b is", a)
		fmt.Println("value of b is", *a)
	}

	// 向函数传递指针参数
	{
		a := 58
		// 58
		fmt.Println("value of a before function call is", a)
		b := &a
		change(b)
		// 55
		fmt.Println("value of a after function call is", a)
	}

	// 不要向函数传递数组的指针，而应该使用切片

	// 特别注意：Go 不支持指针运算
	{
		b := [...]int{109, 110, 111}
		p := &b
		// 注意p++ 是错误的
		//p++
		// [109 110 111]
		fmt.Println(*p)
	}



	fmt.Println("UseCase11Pointer End")
	fmt.Println()
}