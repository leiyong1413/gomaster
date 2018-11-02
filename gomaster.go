// package xxx 每个go文件都应该在开头进行 package xxx 的声明
// 可执行程序的包名应该是main
package main

// import "xxx" 导入xxx 包
// fmt 包 用于打印文本到标准输出
import (
	"github.com/leiyong1413/gomaster/basic"
)

// main 一个特殊函数
// 整个程序都是从main函数开始运行
// main 函数必须放在main包中
func main() {
	basic.UseCase01Introduce()
	basic.UseCase02Variable()
	basic.UseCase03Type()
	basic.UseCase04Constant()
	basic.UseCase05Func()
	basic.UseCase06Package()
	basic.UseCase07ControlStatement()
	basic.UseCase08ArraySlice()
	basic.UseCase09Map()
	basic.UseCase10String()
	basic.UseCase11Pointer()
	basic.UseCase12Struct()
	basic.UseCase13Method()
	basic.UseCase14Interface()
	basic.UseCase15Concurrency()
}
