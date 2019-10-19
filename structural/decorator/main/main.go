package main

import "project/GoDesignPattern/structural/decorator"

func Double(n int) int {
	return 2 * n
}

func main() {
	decoratedFunc := decorator.LogDecorate(Double)
	decoratedFunc(5)
}
