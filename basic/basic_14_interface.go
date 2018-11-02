package basic

import (
	"fmt"
)

// 接口

// 面向对象中接口：接口定义一个对象的行为，做什么的抽象

// 在 Go 语言中，接口就是方法签名（Method Signature）的集合，当一个类型定义了接口中的所有方法，我们称它实现了该接口
// 接口指定了一个类型应该具有的方法，并由该类型决定如何实现这些方法

// 声明和定义
type VowelsFinder interface {
	// 找出所有元音字母
	FindVowels() []rune
}

type MyString string

// 实现接口中的方法
// 而在 Go 中，并不需要这样。如果一个类型包含了接口中声明的所有方法，那么它就隐式地实现了 Go 接口
func (my MyString) FindVowels() []rune  {
	var vowels []rune

	for _,r := range my {
		if r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u' {
			vowels = append(vowels, r)
		}
	}

	return vowels
}

// 接口的实际用途
// 计算工资小列子
type SalaryCalculator interface {
	CalculateSalary() int
}

// 临时工
type Permanent struct {
	empId    int
	basicpay int
	pf       int
}

// 合同工
type Contract struct {
	empId  int
	basicpay int
}

// salary of permanent employee is sum of basic pay and pf
func (p Permanent) CalculateSalary() int {
	return p.basicpay + p.pf
}

// salary of contract employee is the basic pay alone
func (c Contract) CalculateSalary() int {
	return c.basicpay
}

// 计算所有员工的总薪资
func totalExpense(s []SalaryCalculator) int {
	// 遍历接口切片
	expense := 0

	for _,v := range s {
		expense = expense + v.CalculateSalary()
	}

	return expense
}

// 关于空接口
// 没有包含方法的接口称为空接口。空接口表示为 interface{}
// 由于空接口没有方法，因此所有类型都实现了空接口

func describle(i interface{})  {
	fmt.Printf("type = %T, value = %v\n", i, i)
}


// 类型选择
func findType(i interface{})  {
	// 类型断言的语法是 i.(T)，而对于类型选择，类型 T 由关键字 type 代替
	// 类型选择用于将接口的具体类型与很多 case 语句所指定的类型进行比较。它与一般的 switch 语句类似
	// 如果 case 匹配成功，会打印出相应的语句
	switch i.(type) {
	case string:
		fmt.Println("string type")
	case int:
		fmt.Println("int type")
	default:
		fmt.Println("unknown type")
	}
}

// 接口的嵌套
type LeaveCalculator interface {
	CalculateLeavesLeft() int
}

type EmployeeOperations interface {
	// 嵌套接口
	SalaryCalculator
	LeaveCalculator
}

// 接口的零值是nil
// 对于值为 nil 的接口，其底层值（Underlying Value）和具体类型（Concrete Type）都为 nil




func UseCase14Interface()  {
	fmt.Println("UseCase14Interface Start")

	name := MyString("sam anderson")
	// 定义一个接口变量
	var v VowelsFinder
	// 我们把 name 赋值给了 v。由于 MyString 实现了 VowelFinder，因此这是合法的
	v = name
	// v.FindVowels() 调用了 MyString 类型的 FindVowels 方法
	retVowels := v.FindVowels()
	fmt.Println(retVowels)

	// 计算总的支出
	pemp1 := Permanent{1,5000, 20}
	pemp2 := Permanent{1,6000, 30}
	cemp1 := Contract{3, 3000}
	employees := []SalaryCalculator{pemp1,pemp2,cemp1}
	retTotalExpense := totalExpense(employees)
	fmt.Println(retTotalExpense)


	// 关于空接口
	{
		str := "hello world"
		// type = string, value = hello world
		describle(str)

		i := 55
		// type = int, value = 55
		describle(i)

		strt := struct {
			name string
		}{
			name: "Naveen R",
		}
		// type = struct { name string }, value = {Naveen R}
		describle(strt)
	}

	// 类型断言:用于提取接口的底层值
	// 在语法 i.(T) 中，接口 i 的具体类型是 T
	{
		var s interface{} = 56
		// 语法 s.(int) 来提取 s 的底层 int 值
		tmp := s.(int)
		// 56
		fmt.Println(tmp)
		// 如果具体类型不是int，会发生什么？
		var s1 interface{} = "string"
		// 下面语法会报错panic
		//tmp1 := s1.(int)
		// 使用下面语法
		tmp1 , ok := s1.(int)
		if ok {
			fmt.Println(tmp1)
		}else{
			fmt.Println("not int type")
		}
	}

	// 类型选择

	// 特别注意：
	// 指针接受者与值接受者实现接口的不同
	// 值接受者：使用值接受者声明的方法，既可以用值来调用，也能用指针调用，实现的接口也是一样
	// 指针接受者：
	
	// 实现多个接口和接口的嵌套
	// 尽管 Go 语言没有提供继承机制，但可以通过嵌套其他的接口，创建一个新接口


	// 接口的零值是nil
	// 对于值为 nil 的接口，其底层值（Underlying Value）和具体类型（Concrete Type）都为 nil
	var eOps EmployeeOperations
	if eOps == nil {
		// oOps is nil and has type = <nil> value <nil>
		fmt.Printf("oOps is nil and has type = %T value %v\n", eOps, eOps)
		// 注意
		// 对于值为 nil 的接口，由于没有底层值和具体类型，当我们试图调用它的方法时，程序会产生 panic 异常
	} else {

	}



	fmt.Println("UseCase14Interface End")
	fmt.Println()
}