package main

import (
	"fmt"
	"project/GoDesignPattern/creational/singleton"
)

func main() {

	s := singleton.GetInstance()
	s["a"] = "a Value"

	s2 := singleton.GetInstance()

	fmt.Println("s2[a]=", s2["a"])

}
