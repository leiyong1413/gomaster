package basic

import "fmt"

// golang map 键值容器
// map 键的有效性：所有可比较的类型


func UseCase09Map()  {
	fmt.Println("UseCase09Map Start")

	// map 的创建
	// 通过make 函数创建
	// make(map[type of key]type of value)
	personSalary := make(map[string]int)
	// map[]
	// map 的零值是 nil
	// 如果你想添加元素到 nil map 中，会触发运行时 panic。因此 map 必须使用 make 函数初始化
	fmt.Println(personSalary)

	personSalary["steve"] = 12000
	personSalary["jamie"] = 15000
	personSalary["mike"] = 9000
	// personSalary map contents: map[steve:12000 jamie:15000 mike:9000]
	fmt.Println("personSalary map contents:", personSalary)

	{
		var personSalary map[string]int
		if personSalary == nil {
			personSalary = make(map[string]int)
		}
	}

	// 声明的时候初始化map
	{
		personSalary := map[string]int{
			"steve":12000,
			"jamie":15000, // 注意这个逗号
		}
		personSalary["mike"] = 9000
		// personSalary map contents: map[jamie:15000 mike:9000 steve:12000]
		// map 是无序的
		fmt.Println("personSalary map contents:", personSalary)
	}

	// 获取 map 中的元素
	// 获取 map 元素的语法是 map[key] , 如果获取一个不存在的元素，会发生什么呢？map 会返回该元素类型的零值
	// 怎么判断元素是否存在, value, ok := map[key] 如果 ok 是 true，表示 key 存在，key 对应的值就是 value ，反之表示 key 不存在
	{
		personSalary := map[string]int{
			"steve": 12000,
			"jamie": 15000,
		}
		personSalary["mike"] = 9000
		newEmp := "joe"
		value, ok := personSalary[newEmp]
		if ok == true {
			fmt.Println("Salary of", newEmp, "is", value)
		} else {
			fmt.Println(newEmp,"not found")
		}
		// joe not found
	}

	// 遍历 map 中所有的元素需要用 for range 循环
	// 不保证每次执行程序获取的元素顺序相同
	{
		personSalary := map[string]int{
			"steve": 12000,
			"jamie": 15000,
		}
		personSalary["mike"] = 9000
		fmt.Println("All items of a map")
		for key, value := range personSalary {
			fmt.Printf("personSalary[%s] = %d\n", key, value)
		}
		// personSalary[steve] = 12000
		// personSalary[jamie] = 15000
		// personSalary[mike] = 9000
	}

	// 删除 map 中的元素
	//  delete(map, key) 无返回值

	// 获取 map 的长度  len 函数

	// Map 是引用类型
	// 当 map 被赋值为一个新变量的时候，它们指向同一个内部数据结构。因此，改变其中一个变量，就会影响到另一变量
	// 当 map 作为函数参数传递时也会发生同样的情况。函数中对 map 的任何修改，对于外部的调用都是可见的
	{
		personSalary := map[string]int{
			"steve": 12000,
			"jamie": 15000,
		}
		personSalary["mike"] = 9000
		// Original person salary map[steve:12000 jamie:15000 mike:9000]
		fmt.Println("Original person salary", personSalary)
		newPersonSalary := personSalary
		newPersonSalary["mike"] = 18000
		// Person salary changed map[steve:12000 jamie:15000 mike:18000]
		fmt.Println("Person salary changed", personSalary)
	}

	// map 的相等性
	// map 之间不能使用 == 操作符判断，== 只能用来检查 map 是否为 nil
	// 如果判断map是否相等，遍历比较两个map中的每个元素是否相等






	fmt.Println("UseCase09Map End")
	fmt.Println()
}
