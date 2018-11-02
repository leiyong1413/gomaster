package basic

import "fmt"

// 控制语句
// 条件语句
// 	if 条件语句
//  switch 条件语句
// 循环语句

func UseCase07ControlStatement() {
	fmt.Println("UseCase07ControlStatement Start")

	// if 语句
	var condition bool
	if condition {
		fmt.Println("if condition is true show this")
	}

	// if else 语句
	num := 10
	// 求余数
	if num%2 == 0 {
		fmt.Println("the number is even")
	} else {
		fmt.Println("the number is odd")
	}

	// if 语句的另外一种形式
	// if statement; condition {
	// }
	// 它包含一个 statement 可选语句部分，该组件在条件判断之前运行
	{
		if num := 11; num%2 == 0 {
			fmt.Println("the number is even")
		} else {
			fmt.Println("the number is odd")
		}

		// 注意上面的num 的作用域范围只在if else 语句中
	}

	// if else if
	{
		if num := 99; num <= 50 {
			fmt.Println("number is less than or equal to 50")
		} else if num >= 51 && num <= 100 {
			fmt.Println("number is between 51 and 100")
		} else {
			fmt.Println("number is greater than 100")
		}
	}

	// 注意 if else 等语法的格式
	// else 语句应该在 if 语句的大括号 } 之后的同一行中 如果不是，编译器会不通过

	// switch 条件语句
	// 用于将表达式的值与可能匹配的选项列表进行比较，并根据匹配情况执行相应的代码块
	// case 不允许出现重复项
	// switch 和if类似，它包含一个 statement 可选语句部分，该组件在条件判断之前运行
	switch finger := 4; finger {
	case 1:
		fmt.Println("Thumb")
	case 2:
		fmt.Println("Index")
	case 3:
		fmt.Println("Middle")
	case 4:
		fmt.Println("Ring")
	case 5:
		fmt.Println("Pink")
	default: // 默认情况
		fmt.Println("incorrect finger number")
	}

	// 多表达式判断
	// 可以在一个 case 中包含多个表达式
	// 在 case "a","e","i","o","u": 这一行中,列举了所有的元音。只要匹配该项，则将输出 vowel
	switch letter := "i"; letter {
	case "a", "e", "i", "o", "u": // 一个选项多个表达式
		fmt.Println("vowel")
	default:
		fmt.Println("not a vowel")
	}

	// 无表达式的 switch
	// 如果省略表达式，则表示这个 switch 语句等同于 switch true，并且每个 case 表达式都被认定为有效，相应的代码块也会被执行
	number := 75
	switch {
	case number >= 0 && number <= 50:
		fmt.Println("num is greater than 0 and less than 50")
	case number >= 51 && number <= 100:
		fmt.Println("num is greater than 51 and less than 100")
	case number >= 101:
		fmt.Println("num is greater than 100")
	}

	// Fallthrough
	// 在 Go 中，每执行完一个 case 后，会从 switch 语句中跳出来，不再做后续 case 的判断和执行。使用 fallthrough 语句可以在已经执行完成的 case 之后，把控制权转移到下一个 case 的执行代码中
	switch number := 75; true {
	case number < 50:
		fmt.Printf("%d is lesser than 50\n", number)
		fallthrough
	case number < 100:
		fmt.Printf("%d is lesser than 100\n", number)
		fallthrough
	case number < 200:
		fmt.Printf("%d is lesser than 200\n", number)
	}

	// 注意上面如果true 改成 switch number := 75; number { 与后面的case 将报类型不匹配错误，case 选项的值是bool值，所以上面改成了true，而number是int类型 这里需要注意
	// fallthrough 语句应该是 case 子句的最后一个语句。如果它出现在了 case 语句的中间，编译器将会报错：fallthrough statement out of place



	// 循环语句
	// for 是 Go 语言唯一的循环语句
	// for initialisation; condition; post {
	// }

	// 过程：初始化语句只执行一次。循环初始化后，将检查循环条件。如果条件的计算结果为 true ，则 {} 内的循环体将执行，接着执行 post 语句。
	// post 语句将在每次成功循环迭代后执行。在执行 post 语句后，条件将被再次检查。如果为 true, 则循环将继续执行，否则 for 循环将终止

	// 注意三部分都是可选的，即初始化，条件和 post 都是可选的

	for i := 1; i <= 10; i++ {
		fmt.Println("value is", i)
	}

	// i 在for循环作用域中，在循环外不能访问

	// break 跳出整个循环
	for i := 1; i <= 10; i++ {
		if i > 5 {
			break
		}
		fmt.Println("value is", i)
	}

	// continue 跳出当前循环
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("odd value is", i)
	}

	// 无限循环
	// for{
	//
	// }

	fmt.Println("UseCase07ControlStatement End")
	fmt.Println()
}
