package main

import (
	"fmt"
	"math"
)

func main() {
	// var investmentAmount float64 = 2000
	investmentAmount, years := 2000.00, 10.00
	expectedReturnRate := 5.5
	// years := 10.0
	// var futureValue = float64(investmentAmount) * math.Pow((1+expectedReturnRate/100), float64(years))
	futureValue := (investmentAmount) * math.Pow((1+expectedReturnRate/100), (years))
	fmt.Println(futureValue)
}
