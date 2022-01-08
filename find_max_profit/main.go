package main

import (
	"fmt"
)

func findMaxProfit(prices []int) int {
	var lowestPrice int = prices[0]
	var maxProfit int
	var thisProfit int

	for _, v := range prices {
		if v < lowestPrice {
			lowestPrice = v
		}

		if v > lowestPrice {
			thisProfit = v - lowestPrice
		}

		if thisProfit > maxProfit {
			maxProfit = thisProfit
		}
	}

	return maxProfit
}

func main() {
	prices1 := []int{10, 7, 5, 8, 11, 2, 6}
	prices2 := []int{32, 13, 4, 14, 34, 25, 22, 75, 63, 50}
	fmt.Println(findMaxProfit(prices1))
	fmt.Println(findMaxProfit(prices2))
}

/*
3. You’re working on some more stock-prediction software. The function
you’re writing accepts an array of predicted prices for a particular stock
over the course of time.
For example, this array of seven prices:
[10, 7, 5, 8, 11, 2, 6]
predicts that a given stock will have these prices over the next seven days.
(On Day 1, the stock will close at $10; on Day 2, the stock will close at
$7; and so on.)

Your function should calculate the greatest profit that could be made
from a single “buy” transaction followed by a single “sell” transaction.
In the previous example, the most money could be made if we bought the
stock when it was worth $5 and sold it when it was worth $11. This yields
a profit of $6 per share.

Note that we could make even more money if we buy and sell multiple
times, but for now, this function focuses on the most profit that could be
made from just one purchase followed by one sale.

Now, we could use nested loops to find the profit of every possible buyand-
sell combination. However, this would be O(N^2) and too slow for our
hotshot trading platform. Your job is to optimize the code so that the
function clocks in at just O(N).

set profit = 0
set lowest = slice[0]

if lowest > currentNumber
  lowest = currentNumber
else
  calculate profit
    if newProfit > profit
      newProfit = profit

return
- greatest difference between two numbers, in order
*/
