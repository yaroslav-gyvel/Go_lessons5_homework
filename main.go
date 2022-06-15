package main

import (
	"fmt"
	"math"
)

func main() {

	const (
		priceOneApple = 5.99
		priceOnePear  = 7
		haveMoney     = 23
		countApples   = 9
		countPears    = 8
	)

	var (
		costApples float64
		costPears  float64
		sumMoney   float64
	)

	fmt.Println("Price one apple", priceOneApple, "uah")
	fmt.Println("Price one pear", priceOnePear, "uah")
	fmt.Println("We have money", haveMoney, "uah")
	fmt.Println("")
	costApples = countApples * priceOneApple
	costPears = countPears * priceOnePear
	sumMoney = costApples + costPears
	fmt.Println("How much money do you have to spend to buy 9 apples and 8 pears?", sumMoney, "uah")
	fmt.Println("")
	buyPears := float64(haveMoney) / float64(priceOnePear)
	int, frac := math.Modf(buyPears)
	_ = frac
	fmt.Println("How many pears can we buy?", int, "pieces")
	fmt.Println("")
	buyApples := float64(haveMoney) / float64(priceOneApple)
	i, f := math.Modf(buyApples)
	_ = f
	fmt.Println("How many apples can we buy?", i, "pieces")
	fmt.Println("")
	sumMoney = 2*priceOneApple + 2*priceOnePear
	fmt.Println("Can we buy 2 pears and 2 apples?", sumMoney < float64(haveMoney))
}
