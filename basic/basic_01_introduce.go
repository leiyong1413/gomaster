// package xxx 每个go文件都应该在开头进行 package xxx 的声明
package basic

import "fmt"

func useCase01Reference() {
	var contents = `
		Introduce
	`
	println(contents)
}

func UseCase01Introduce() {
	fmt.Println("UseCase01Introduce Start")

	fmt.Println("UseCase01Introduce End")
	fmt.Println()
}
