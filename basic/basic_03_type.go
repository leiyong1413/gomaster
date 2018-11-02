package basic

import (
	"fmt"
	"unsafe"
)

// 类型
// 基本类型：
// 	bool
// 	数字类型(int8 int16 int32 int64 int uint8 uint16 uint32 uint64 uint float32 float64 complex64 complex128 byte rune)
//  string
func UseCase03Type() {
	fmt.Println("UseCase03Type Start")
	// bool 类型 布尔值 值为true或false
	{
		a := true
		b := false
		fmt.Println("a", a, "b", b)
		c := a && b
		fmt.Println("c", c)
		d := a || b
		fmt.Println("d", d)
	}

	// 有符号整型
	// int8：表示 8 位有符号整型
	// 大小：8 位
	// 范围：-128～127
	//
	// int16：表示 16 位有符号整型
	// 大小：16 位
	// 范围：-32768～32767
	//
	// int32：表示 32 位有符号整型
	// 大小：32 位
	// 范围：-2147483648～2147483647
	//
	// int64：表示 64 位有符号整型
	// 大小：64 位
	// 范围：-9223372036854775808～9223372036854775807

	// int：根据不同的底层平台（Underlying Platform），表示 32 或 64 位整型。除非对整型的大小有特定的需求，否则你通常应该使用 int 表示整型。
	// 大小：在 32 位系统下是 32 位，而在 64 位系统下是 64 位。
	// 范围：在 32 位系统下是 -2147483648～2147483647，而在 64 位系统是 -9223372036854775808～9223372036854775807

	// printf 中使用%T 打印变量类型
	// Go 的 unsafe 包提供了一个 Sizeof 函数，该函数接收变量并返回它的字节大小
	{
		// 1个字节
		var a int8
		fmt.Printf("value of a is %d, type of a is %T, size of a is %d\n", a, a, unsafe.Sizeof(a))
	}

	// 无符号整型
	// uint8：表示 8 位无符号整型
	// 大小：8 位
	// 范围：0～255
	//
	// uint16：表示 16 位无符号整型
	// 大小：16 位
	// 范围：0～65535
	//
	// uint32：表示 32 位无符号整型
	// 大小：32 位
	// 范围：0～4294967295
	//
	// uint64：表示 64 位无符号整型
	// 大小：64 位
	// 范围：0～18446744073709551615
	//
	// uint：根据不同的底层平台，表示 32 或 64 位无符号整型。
	// 大小：在 32 位系统下是 32 位，而在 64 位系统下是 64 位。
	// 范围：在 32 位系统下是 0～4294967295，而在 64 位系统是 0～18446744073709551615

	// 浮点型
	// float32：32 位浮点数
	// float64：64 位浮点数
	{
		a, b := 5.67, 8.90
		// float64 是浮点数的默认类型
		fmt.Printf("type of a %T b %T\n", a, b)
		sum := a + b
		diff := a - b
		fmt.Println("sum is", sum, "diff is", diff)
	}

	// 复数类型
	// complex64：实部和虚部都是 float32 类型的的复数。
	// complex128：实部和虚部都是 float64 类型的的复数
	{
		// 创建复数
		c := 6 + 7i
		// complex128 是复数的默认类型
		fmt.Printf("type of c %T\n", c)
	}

	// 其他数字类型
	// byte 是 uint8 的别名
	// rune 是 int32 的别名
	// 在字符串中使用
	{

	}

	// string 类型
	// 在 Golang 中，字符串是字节的集合， 注意是字节，本质是字节的切片，详见 basic_10_string
	{
		first := "lei"
		last := "yong"
		// + 操作符可以用于拼接字符串
		fullName := first + " " + last
		fmt.Println("my name is", fullName)
	}

	// 类型转换
	// Go 有着非常严格的强类型特征。Go 没有自动类型提升或类型转换
	// 即使byte是uint8的别名也是一样，是不同的类型
	{
		var a int8 = 5     // int
		var b byte = 1     // byte
		sum := a + int8(b) // mismatched types int8 and byte
		fmt.Printf("sum is %d and type is %T\n", sum, sum)
	}

	fmt.Println("UseCase03Type End")
	fmt.Println()
}
