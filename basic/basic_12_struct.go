package basic

import "fmt"

// 结构体
// 结构体是用户定义的类型，表示若干个字段（Field）的集合

// 结构体的声明
// 声明了一个结构体类型 Employee
// 它有 firstName、lastName 和 age 三个字段
type Employee struct {
	firstName string
	lastName string
	age int
}

type ReEmployee struct {
	firstName, lastName string // 通过把相同类型的字段声明在同一行，结构体可以变得更加紧凑
	age int
}

// 声明结构体时也可以不用声明一个新类型
//  匿名结构体
var employee struct{
	firstName, lastName string
	age int
}

type Address struct {
	city, state string
}

type Person struct {
	name string
	age int
	Address
}

func UseCase12Struct()  {
	fmt.Println("UseCase12Struct Start")

	// 创建一个结构变量
	// 字段名的顺序不一定要与声明结构体类型时的顺序相同
	emp1 := Employee{
		firstName:"Sam",
		age:25,
		lastName:"Anderson",
	}

	// 省略了字段名。在这种情况下，就需要保证字段名的顺序与声明结构体时的顺序相同
	emp2 := Employee{"Thomas","Paul", 30}
	// employee 1 {Sam Anderson 25}
	fmt.Println("employee 1", emp1)
	// employee 2 {Thomas Paul 30}
	fmt.Println("employee 2", emp2)

	// 创建匿名结构体
	// 之所以称这种结构体是匿名的，是因为它只是创建一个新的结构体变量 em3，而没有定义任何结构体类型
	emp3 := struct {
		firstName, lastName string
		age, salary int
	}{
		firstName:"Andreah",
		lastName:"Nikola",
		age:31,
		salary:5000,
	}
	// employee 3 {Andreah Nikola 31 5000}
	fmt.Println("employee 3", emp3)

	// 结构体的零值
	// 当定义好的结构体并没有被显式地初始化时，该结构体的字段将默认赋为零值
	//  string 的零值 ""
	var emp4 Employee
	// employee 4 {  0}
	fmt.Println("employee 4", emp4)

	// 可以为某些字段指定初始值，而忽略其他字段。这样，忽略的字段名会赋值为零值
	emp5 := Employee{
		firstName: "John",
		lastName:  "Paul",
	}
	// Employee 5 {John Paul 0}
	fmt.Println("Employee 5", emp5)

	// 访问结构体的字段,点号操作符 . 用于访问结构体的字段
	fmt.Println("first name", emp5.firstName)
	fmt.Println("last name", emp5.lastName)

	// 结构体的指针,指向结构体的指针
	emp6 := &Employee{"sam","anderson",30}
	fmt.Println("first name", (*emp6).firstName)
	// Go 语言允许我们在访问 firstName 字段时，可以使用 emp8.firstName 来代替显式的解引用 (*emp8).firstName
	fmt.Println("first name", emp6.firstName)

	// 匿名字段
	// 创建结构体时，字段可以只有类型，而没有字段名。这样的字段称为匿名字段
	// 虽然匿名字段没有名称，但其实匿名字段的名称就默认为它的类型
	// 结构体的字段有可能也是一个结构体, 这样的结构体称为嵌套结构体
	// 提升字段:如果是结构体中有匿名的结构体类型字段，则该匿名结构体里的字段就称为提升字段。这是因为提升字段就像是属于外部结构体一样，可以用外部结构体直接访问
	p1 := Person{
		name:"p1",
		age:30,
		Address:Address{
			city:"chengdu",
			state:"sichuan",
		},
	}
	// p1 {p1 30 {chengdu sichuan}}
	fmt.Println("p1",p1)

	// 提升字段访问
	fmt.Println("name", p1.name)
	fmt.Println("city",p1.city)


	// 关于导出：如果结构体名称以大写字母开头，则它是其他包可以访问的导出类型
	// 如果结构体里的字段首字母大写，它也能被其他包访问到


	// 关于相等：
	// 结构体是值类型。如果它的每一个字段都是可比较的，则该结构体也是可比较的。如果两个结构体变量的对应字段相等，则这两个变量也是相等的
	// 值类型


	fmt.Println("UseCase12Struct End")
	fmt.Println()
}




