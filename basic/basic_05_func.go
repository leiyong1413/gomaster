package basic

import "fmt"

// 对可变参数函数的理解是一个小难点

// 函数

// 函数声明的语法
// func functionName(parameterName type) returnType {
//    // 函数体（具体实现的功能）
// }
// 任意多个参数采用类似 (parameter1 type, parameter2 type) 即(参数1 参数1的类型,参数2 参数2的类型)的形式指定

// 函数中的参数列表和返回值并非是必须的
func functionName() {
	// 译注: 表示这个函数不需要输入参数，且没有返回值
}

// 计算商品价格
func calculateBill(price int, number int) int {
	totalPrice := price * number
	return totalPrice
}

// 如果有连续若干个参数，它们的类型一致，那么我们无须一一罗列，只需在最后一个参数后添加该类型
func reCalculateBill(price, number int) int {
	totalPrice := price * number
	return totalPrice
}

// 多返回值
// 如果一个函数有多个返回值，那么这些返回值必须用 ( 和 ) 括起来
func calculateRectProps(length, width float64) (float64, float64) {
	area := length * width
	perimeter := (length + width) * 2
	return area, perimeter
}

// 命名返回值
// 从函数中可以返回一个命名值。一旦命名了返回值，可以认为这些值在函数第一行就被声明为变量了
func reCalculateRectProps(length, width float64) (area float64, perimeter float64) {
	area = length * width
	perimeter = (length + width) * 2

	// 不需要明确指定返回值，默认返回 area, perimeter 的值
	return
}

func find(num int, nums ...int)  {
	// type of nums is []int 可以知道是一个切片
	fmt.Printf("type of nums is %T\n", nums)

	found := false
	for _, v := range nums {
		if v == num {
			fmt.Println(num, "found in ", nums)
			found = true
			break
		}
	}

	if !found {
		fmt.Println(num , "not found in", nums)
	}
	fmt.Println()
}

func UseCase05Func() {
	fmt.Println("UseCase05Func Start")

	price, number := 99, 6
	totalPrice := reCalculateBill(price, number)
	fmt.Println("total price is", totalPrice)

	length, width := 40.0, 50.7
	area, perimeter := calculateRectProps(length, width)
	fmt.Println("area is", area, "and perimeter  is", perimeter)

	// 空白符号
	// _ 在 Go 中被用作空白符，可以用作表示任何类型的任何值
	// 空白符 _ 用来跳过不要的计算结果
	onlyArea, _ := reCalculateRectProps(length, width)
	fmt.Println("area is", onlyArea)

	// 可变参数的函数
	// 如果函数最后一个参数被记作 ...T ，这时函数可以接受任意个 T 类型参数作为最后一个参数
	// 可变参数函数的工作原理:是把可变参数转换为一个新的切片
	find(89, 89, 90,95)

	// 给可变参数函数传入切片, 是不行的？ 为什么？
	// 看find定义知道接受的类型是int类型，可变参数参数会被转换为 int 类型切片然后在传入 find 函数中。
	// 但是在这里 nums 已经是一个 int 类型切片，编译器试图在 nums 基础上再创建一个切片,所以错误

	// 如何修改：有一个可以直接将切片传入可变参数函数的语法糖，你可以在在切片后加上 ... 后缀。
	// 如果这样做，切片将直接传入函数，不再创建新的切片
	nums := []int{89, 90,95}
	find(89,nums...)

	fmt.Println("UseCase05Func End")
	fmt.Println()
}
