package main

import "fmt"
import "github.com/leiyong1413/gomaster/basic/gettingstarted"

type UseCaseType int

// 用列枚举
const (
	BasicGettingStarted UseCaseType = iota
	BasicPackage
)

// 实现一个简单工厂
type UseCase interface {
	Reference()
	RunUseCase()
}

// 简单工厂
func NewUseCase(useCaseType UseCaseType) UseCase {
	switch useCaseType {
	case BasicGettingStarted:
		return &gettingstarted.GettingStarted{}
	case BasicPackage:
	default:
	}
	return nil
}

func main() {
	fmt.Println("hello golang!!!")
	useCase := NewUseCase(BasicGettingStarted)
	useCase.RunUseCase()
}
