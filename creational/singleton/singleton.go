package singleton

import (
	"fmt"
	"sync"
)

//创建一个类型为map的单例
type singleton map[string]string

var (
	once     sync.Once
	instance singleton
)

func GetInstance() singleton {
	once.Do(func() {
		instance = make(map[string]string)
		fmt.Println("make object")
	})
	return instance
}
