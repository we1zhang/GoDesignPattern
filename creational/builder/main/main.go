package main

import (
	"fmt"
	car "project/GoDesignPattern/creational/builder"
)

func main() {


	car := (&car.CarBuilder{}).SetColor(car.BLUE).
		SetSpeed(car.MPH).
		SetWheels(car.SPORTS_WHEELS).
		Create()

	fmt.Println("car:", *car)

}
