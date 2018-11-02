package basic

import "fmt"

// 字符串

// Go 语言中的字符串是一个字节切片。把内容放在双引号""之间，我们可以创建一个字符串
// Go 中的字符串是兼容 Unicode 编码的，并且使用 UTF-8 进行编码

func printBytes(s string)  {
	// len(s) 返回字符串中字节的数量
	for i := 0; i < len(s);i++  {
		// %x 格式限定符用于指定 16 进制编码
		fmt.Printf("%x", s[i])
	}
}

// 打印字符时，假定每个字符只占用一个字节
func printChars(s string) {
	for i:= 0; i < len(s); i++ {
		// %c 格式限定符用于打印字符串的字符
		fmt.Printf("%c ",s[i])
	}
}

func printCharsRune(s string) {
	runes := []rune(s)
	for i:= 0; i < len(runes); i++ {
		fmt.Printf("%c ",runes[i])
	}
}

func UseCase10String()  {
	fmt.Println("UseCase10String Start")

	name := "hello world"
	fmt.Println(name)

	name = "Señor"
	printBytes(name)
	// 单独获取字符串的每一个字节
	
	// 在 UTF-8 编码中，一个代码点可能会占用超过一个字节的空间。那么我们该怎么办呢？rune 能帮我们解决这个难题
	// rune 是 Go 语言的内建类型，它也是 int32 的别称。在 Go 语言中，rune 表示一个代码点。代码点无论占用多少个字节，都可以用一个 rune 来表示
	printCharsRune(name)

	// 字符串的 for range 循环
	// 但是 Go 给我们提供了一种更简单的方法来做到这一点：使用 for range 循环
	for index, rune := range name {
		// 循环返回的是是当前 rune 的字节位置
		fmt.Printf("%c starts at byte %d\n", rune, index)
	}

	// 用字节切片构造字符串
	byteSlice := []byte{0x43, 0x61, 0x66, 0xC3, 0xA9}
	str := string(byteSlice)
	fmt.Println(str)

	// 用 rune 切片构造字符串
	runeSlice := []rune{0x0053, 0x0065, 0x00f1, 0x006f, 0x0072}
	runeStr := string(runeSlice)
	fmt.Println(runeStr)

	// 字符串的长度
	// func RuneCountInString(s string) (n int) 方法用来获取字符串的长度 返回字符串中的 rune 的数量

	// 字符串是不可变的
	// Go 中的字符串是不可变的。一旦一个字符串被创建，那么它将无法被修改
	// 为了修改字符串，可以把字符串转化为一个 rune 切片。然后这个切片可以进行任何想要的改变，然后再转化为一个字符串
	tmpRune := []rune("hello")
	tmpRune[0] = 'a'
	// aello
	fmt.Println(string(tmpRune))

	fmt.Println("UseCase10String End")
	fmt.Println()
}