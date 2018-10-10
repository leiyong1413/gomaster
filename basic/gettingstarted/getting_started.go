package gettingstarted

import "fmt"

type GettingStarted struct {
}

func (*GettingStarted) Reference() {
	// `` 可以多行，但是多行换行符号go会修改为\n 不是编译器的\r\n
	reference := `
	
	`
	fmt.Println(reference)
}

func (*GettingStarted) RunUseCase() {
	fmt.Println("run use case: getting started")
}
