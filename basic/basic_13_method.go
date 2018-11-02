package basic

import "fmt"

// 方法
// 带有接收器的函数
// 接收器：可以是结构体类型或非结构体类型
// 接收器，可在方法内部访问

// 定义：
// func (t Type) methodName(parameter list) {
//
// }
// 接收器类型：Type

// 为什么要定义方法？
// Go 不是纯粹的面向对象编程语言，而且Go不支持类。因此，基于类型的方法是一种实现和类相似行为的途径
// 相同的名字的方法可以定义在不同的类型上，而相同名字的函数是不被允许的

// 关于指针接收器和值接收器？
// 在指针接收器的方法内部的改变对于调用者是可见的，然而值接收器的情况不是这样的

type MAddress struct {
	city, state string
}

func (a MAddress) fullAddress()  {
	fmt.Printf("full address:%s, %s\n", a.city, a.state)
}

type MEmpolyee struct {
	name string
	salary int
	currency string
	MAddress
}

// 定义一个方法，打印员工信息
func (e MEmpolyee) displayEmpolyeeInfo()  {
	fmt.Printf("salary of %s is %s%d\n",e.name, e.currency, e.salary)
}

// 使用值接收器，改变名字
func (e MEmpolyee) changeName(newName string)  {
	e.name = newName
}

// 使用指针接收器，改变薪资
func (e *MEmpolyee) changeSalary(newSalary int)  {
	(*e).salary = newSalary
	e.salary = newSalary
}

// 非结构类型的方法
type myInt int

// 定义方法
func (a myInt)add(b myInt) myInt  {
	return a + b
}

func UseCase13Method()  {
	fmt.Println("UseCase13Method Start")

	emp1 := MEmpolyee{
		name:"sam adolf",
		salary:5000,
		currency:"$",
		MAddress:MAddress{
			city:"los angles",
			state:"california",
		},
	}
	// 方法的调用
	emp1.displayEmpolyeeInfo()

	// 关于可见性：
	// 在指针接收器的方法内部的改变对于调用者是可见的，然而值接收器的情况不是这样的
	emp1.changeName("new name")
	// sam adolf 可见名字未改变
	fmt.Println("new name", emp1.name)
	//
	(&emp1).changeSalary(6000)
	// 理论上应该使用上面的调用方式，但是Go语言让我们可以直接使用 emp1.changeSalary(6000)
	emp1.changeSalary(6000)
	// 6000 可见薪资已经改变
	fmt.Println("new salary", emp1.salary)

	// 关于什么时候使用指针接收器什么时候使用值接收器 ？
	// 一般来说，指针接收器可以使用在：对方法内部的接收器所做的改变应该对调用者可见时。
	//指针接收器也可以被使用在如下场景：当拷贝一个结构体的代价过于昂贵时


	// 属于结构体的匿名字段的方法可以被直接调用，就好像这些方法是属于定义了匿名字段的结构体一样
	emp1.fullAddress()

	// 关于接收器：
	// 当一个方法有一个值接收器，它可以接受值接收器和指针接收器， go 语言会做相应的转换如：r.method() 解释为(*r).method()
	// 当一个方法有一个指针接收器，它可以接受使用值接收器和指针接收器,go 语言会做相应的转换如：r.method() 解释为(&r).method()

	// 但是注意：如果是函数的参数则不同，函数使用指针参数只接受指针， 函数使用值参数，它只能接受一个值参数



	// 非结构体上的方法，接收器可以是结构体类型或非结构体类型,如常用的类型如int byte等可以自定义自己的类型
	// 为了在一个类型上定义一个方法，方法的接收器类型定义和方法的定义应该在同一个包中

	fmt.Println("UseCase13Method End")
	fmt.Println()
}