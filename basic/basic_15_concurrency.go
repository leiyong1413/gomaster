package basic

import (
	"fmt"
	"time"
)

// Go 是并发式语言，而不是并行式语言

// 并发是什么?
// 并发是指立即处理多个任务的能力

// 并行：指同时处理多个任务

// 举列子：假如这个人在慢跑时，还在用他的 iPod 听着音乐。在这里，他是在跑步的同时听音乐，也就是同时处理多个任务。这称之为并行
// 想象一个人正在跑步。假如在他晨跑时，鞋带突然松了。于是他停下来，系一下鞋带，接下来继续跑。这个例子就是典型的并发

// 并行不一定会加快运行速度，因为并行运行的组件之间可能需要相互通信
// Go 对并发的支持:Go 编程语言原生支持并发。Go 使用 Go 协程（Goroutine） 和信道（Channel）来处理并发


// Go 协程是什么？
// Go 协程是与其他函数或方法一起并发运行的函数或方法。Go 协程可以看作是轻量级线程，与线程相比，创建一个 Go 协程的成本很小
// Go 协程相比于线程的优势:
// 1.相比线程而言，Go 协程的成本极低
// 2.Go 协程会复用（Multiplex）数量更少的 OS 线程
// 3.Go 协程使用信道（Channel）来进行通信

// 如何启动一个 Go 协程？
// 使用关键字go
// 注意：main函数会运行在一个特有的go协程上，它称为 Go 主协程
// 注意调用关系：
// 1.启动一个新的协程时，协程的调用会立即返回。与函数不同，程序控制不会去等待 Go 协程执行完毕。在调用 Go 协程之后，程序控制会立即返回到代码的下一行，忽略该协程的任何返回值
// 2.如果希望运行其他 Go 协程，Go 主协程必须继续运行着。如果 Go 主协程终止，则程序终止，于是其他 Go 协程也不会继续运行，不会等待其他其他协程结束

// 信道？
//  Go 协程之间通信的管道
// 信道的声明：所有信道都关联了一个类型。信道只能运输这种类型的数据，而运输其他类型的数据都是非法的
// chan T:表示 T 类型的信道, 信道的零值是nil， 信道使用make函数定义

// 注意 发送和接受默认是阻塞的,
// 当把数据发送到信道时，程序控制会在发送数据的语句处发生阻塞，直到有其它 Go 协程从信道读取到数据，
// 才会解除阻塞。与此类似，当读取信道的数据时，如果没有其它的协程把数据写入到这个信道，
// 那么读取过程就会一直阻塞着
// 这样做的目的是：帮助 Go 协程之间进行高效的通信，而不用像其他语言那样显式的锁和条件变量


// 缓存信道

func hello(done chan bool)  {
	fmt.Println("hello go routine is going to sleep")
	time.Sleep(4 * time.Second)
	fmt.Println("hello go routine awake and going to write to done")
	done <- true
}

func producer(data chan int)  {
	for i := 0; i < 10; i++ {
		data <- i
	}
	close(data)
}

func producer1(data chan int)  {
	for i:= 0; i < 10; i++ {
		data <- i
	}
	close(data)
}

func UseCase15Concurrency()  {
	fmt.Println("UseCase15Concurrency Start")

	// 定义一个信道 传输int 类型数据
	var a chan int
	if a == nil {
		a = make(chan int)
	}
	// chan int
	fmt.Printf("type of a is %T\n",a)
	// 使用简短方式：
	b := make(chan int)
	// chan int
	fmt.Printf("type of b is %T\n",b)

	// 信道的数据的收发？
	// data := <- a 读取信道a
	// a <- data 写入信道a

	// 注意 发送和接受默认是阻塞的
	//
	done := make(chan bool)
	fmt.Println("Main going to call hello go goroutine")
	go hello(done)
	// 会阻塞，主协程会在 <-done 这一行发生阻塞，等待来自信道 done 的数据
	<- done
	fmt.Println("main received data")

	// 死锁
	// 使用信道需要考虑的一个重点是死锁
	// 当 Go 协程给一个信道发送数据时，照理说会有其他 Go 协程来接收数据。如果没有的话，程序就会在运行时触发 panic，形成死锁
	// 同理，当有 Go 协程等着从一个信道接收数据时，我们期望其他的 Go 协程会向该信道写入数据，要不然程序就会触发 panic
	// chTest := make(chan int)
	// main received data
	// fatal error: all goroutines are asleep - deadlock!
	// chTest <- 5

	// 单向信道：信道只能发送或者接收数据
	// 默认的信道是双向的，即通过信道既能发送数据，又能接收数据
	// 只发送： xxx chan<- int
	// 只接收： xxx <-chan int
	// 注意：把一个双向信道转换成唯送信道或者唯收（Receive Only）信道都是行得通的，但是反过来就不行


	// 关闭信道
	// 数据的发送方可以关闭信道，会通知接收方这个信道不再有数据发送过来
	// close(chan)
	// 接收方如果判断：v, ok := <- ch
	// 如果ok等于false，说明我们试图读取一个关闭的通道，从关闭的信道读取到的值会是该信道类型的零值

	ch := make(chan int)
	go producer(ch)
	// 遍历获取数据
	for {
		v, ok := <-ch
		if ok == false {
			break
		}
		fmt.Println("Received ", v, ok)
	}

	// 使用for range 遍历，for range 循环从信道 ch 接收数据，直到该信道关闭，一旦关闭循环会自动结束
	// 如上面的列子可以简写如下：
	// 重写，提高了代码的质量和可读性
	ch1 := make(chan int)
	go producer1(ch1)
	for v := range ch1 {
		fmt.Println("Received", v)
	}

	fmt.Println("UseCase15Concurrency End")
	fmt.Println()
}



