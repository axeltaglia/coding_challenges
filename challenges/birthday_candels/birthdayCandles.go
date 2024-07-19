/*
You are in charge of the cake for a child's birthday. You have decided the cake will have one candle for each year of their total age. They will only be able to blow out the tallest of the candles. Count how many candles are tallest.

Example

You are in charge of the cake for a child’s birthday. You have decided the cake will have one candle for each year of their total age. They will only be able to blow out the tallest of the candles. Count how many candles are tallest. Count how many candles are tallest.

function birthdayCakeCandles(candles) {

}

Candles represents the sample case array [3, 2 , 1, 3].
Answer: 2
*/

package main

import "fmt"

func main() {
	candles := []int{1, 3, 1, 3}
	result := birthdayCakeCandles(candles)
	fmt.Println(result)
}

func birthdayCakeCandles(candles []int) int {
	longest := 0
	amount := 1

	for _, candle := range candles {
		if candle > longest {
			longest = candle
			amount = 1
		} else if candle == longest {
			amount++
		}
	}

	return amount
}
