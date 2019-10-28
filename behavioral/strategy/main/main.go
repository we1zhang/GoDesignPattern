package main

import (
	"fmt"
	strategy2 "project/GoDesignPattern/behavioral/strategy"
)

func main() {
	var strategy strategy2.TravelStrategy
	way := "ship" //car

	switch way {
	case "car":
		strategy = &strategy2.CarStrategy{}
	case "bike":
		strategy = &strategy2.BikeStrategy{}
	case "ship":
		strategy = &strategy2.ShipStrategy{}
	default:
		return
	}

	fmt.Println(strategy.Travel())
}
