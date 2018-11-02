package basic

import "fmt"

// 对切片的理解是一个小难点


// 数组和切片

// 数组是同一类型元素的集合
// 一个数组的表示形式为 [n]T
// n 表示数组中元素的数量，T 代表每个元素的类型
// 元素的数量 n 也是该类型的一部分

func changeLocal(num [5]int)  {
	num[0] = 55
	fmt.Println("array inside function", num)
}


// 切片
// 切片是由数组建立的一种方便、灵活且功能强大的包装
// 切片本身不拥有任何数据, 只是对现有数组的引用

func UseCase08ArraySlice()  {
	fmt.Println("UseCase08ArraySlice Start")

	var a [3]int
	// 数组中的所有元素都被自动赋值为数组类型的零值 [0 0 0]
	fmt.Println("array a is", a)
	// 数组的索引从 0 开始到 length - 1 结束
	for index:= 0; index < len(a); index++{
		fmt.Println("element", index, "value is", a[index])
	}

	//  简略声明
	{
		a := [3]int{12, 78, 50}
		// [12 78 50]
		fmt.Println(a)

		// 在简略声明中，不需要将数组中所有的元素赋值, 未赋值的元素自动赋零值
		b := [3]int{12}
		// [12 0 0]
		fmt.Println(b)

		// 你甚至可以忽略声明数组的长度，并用 ... 代替，让编译器为你自动计算长度
		c := [...]int{12, 13, 14, 15}
		// [12 13 14 15]
		fmt.Println(c)

		// 数组的大小是类型的一部分。因此 [5]int 和 [25]int 是不同类型
		// 数组不能调整大小，不要担心这个限制,因为 slices 的存在能解决这个问题
	}

	// 数组是值类型
	// Go 中的数组是值类型而不是引用类型。这意味着当数组赋值给一个新的变量时，该变量会得到一个原始数组的一个副本,如果对新变量进行更改，则不会影响原始数组
	{
		a := [...]string{"USA", "China", "India", "Germany", "France"}
		b := a // copy of a
		b[0] = "Singapore"
		// a is [USA China India Germany France]
		fmt.Println("a is", a)
		// b is [Singapore China India Germany France]
		fmt.Println("b is", b)
	}

	// 当数组作为参数传递给函数时，它们是按值传递，而原始数组保持不变
	{
		num := [...]int{5, 6,7,8,9}
		// before passing to function [5 6 7 8 9]
		fmt.Println("before passing to function", num)
		// array inside function [55 6 7 8 9]
		changeLocal(num)
		// after passing to function [5 6 7 8 9]
		fmt.Println("after passing to function", num)
	}

	// 数组的长度
	// 将数组作为参数传递给len函数，得到数组的长度
	{
		a := [...]float64{67.7, 89.8, 21, 78}
		// length of a is 4
		fmt.Println("length of a is", len(a))
	}

	// 数组的迭代
	{
		// 使用for
		a := [...]float64{67.7, 89.8, 21, 78}
		for i := 0; i < len(a); i++ {
			fmt.Printf("%d th element of a is %.2f\n", i, a[i])
		}

		fmt.Println()

		// 通过使用 for 循环的 range 方法来遍历数组
		// range 返回索引和该索引处的值
		sum := 0.0
		//range returns both the index and value
		for i, v := range a{
			fmt.Printf("%d the element of a is %.2f\n", i, v)
			sum += v
		}
		fmt.Println("\n sum of all elements of a", sum)
		fmt.Println()

		// 使用空白符，可忽略索引或该索引处的值
		for _, v := range a{
			fmt.Printf("the element of a is %.2f\n", v)
		}
	}


	// 创建切片
	// 带有 T 类型元素的切片由 []T 表示
	{
		a := [5]int{76, 77, 78, 79, 80}
		// 使用语法 a[start:end] 创建一个从 a 数组索引 start 开始到 end - 1 结束的切片
		var slice []int = a[1:4]
		// [77 78 79]
		fmt.Println(slice)
	}

	// 另外一种创建切片的方式
	{
		// creates and array and returns a slice reference
		s := []int{6,7,8}
		fmt.Println(s)
	}

	// 切片的修改
	// 对切片所做的任何修改都会反映在底层数组中
	// 当多个切片共用相同的底层数组时，每个切片所做的更改将反映在数组中
	{
		orgArray := [...]int{57, 89, 90, 82, 100, 78, 67, 69, 59}
		slice := orgArray[2:5]
		// array before [57 89 90 82 100 78 67 69 59]
		fmt.Println("array before", orgArray)
		for i, _ := range slice {
			slice[i]++
		}
		// array after [57 89 91 83 101 78 67 69 59]
		fmt.Println("array after", orgArray)

		// 开始和结束的默认值分别为 0 和 len (orgArray)
		orgSlice1 := orgArray[:]
		orgSlice2 := orgArray[:]
		// [57 89 91 83 101 78 67 69 59]
		fmt.Println(orgArray)
		orgSlice1[0] = 100
		// [100 89 91 83 101 78 67 69 59]
		fmt.Println(orgArray)
		orgSlice2[1] = 101
		// [100 101 91 83 101 78 67 69 59]
		fmt.Println(orgArray)
	}

	// 切片的长度和容量
	// 长度：切片的长度是切片中的元素数
	// 容量：从创建切片索引开始的底层数组中元素数

	// 使用make函数创建切片
	// func make（[]T，len，cap）[]T 通过传递类型，长度和容量来创建切片。容量是可选参数, 默认值为切片长度。make 函数创建一个数组，并返回引用该数组的切片
	i := make([]int, 5, 5)
	// [0 0 0 0 0]
	fmt.Println(i)

	// 追加切片元素
	// func append（s[]T，x ... T）[]T
	// 这是一个可变函数：x ... T 在函数定义中表示该函数接受参数 x 的个数是可变的

	fmt.Println("UseCase08ArraySlice End")
	fmt.Println()
}