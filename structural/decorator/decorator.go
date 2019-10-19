package decorator

import "log"

type Object func(int) int

func LogDecorate(fn Object) Object {
	return func(n int) int {
		log.Println("starting decorate,n=", n)
		result := fn(n)
		log.Println("end decorate,result=", result)
		return result
	}
}
