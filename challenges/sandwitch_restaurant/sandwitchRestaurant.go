// To execute Go code, please declare a func main() in a package "main"

// Jim owns a sandwich restaurant and he manages it in his own way. While in a normal restaurant, a customer is served by following the first-come, first-served rule. He wants to calculate the average wait time of each of his customers.
// Structure of the data is: [(arrival time, length of time to create a sandwich)…]
// Example: [(0, 3), (1, 9), (2, 6)] The average waiting time is 10 minutes
// 1st Customer Takes 3 minutes. Chef Start time: 0 minutes waits: 3 minutes 3 - 0 + 0 = 3
// 2nd Customer waits 11 minutes Chef Start time: 3 minutes waits: 11 minutes. 9 - 1 + 3 = 11
// 3rd Customer waits 6 - 2 + 11
// Chef Start Time.     Chef End Time
//. 0                   3
//. 3                   12
//. 12                  18
//  50                  51

// 3. - 0 = 3 - 0 + 0
// 12 - 1 = 11 = 9 -1 + 3 = takeTime - startTime + acum
// 18 -2 = 16 = 6 - 2 + 11
//(3 + 11 + 16)/3 = 30/3 = 10

// Chef EndTime - customer arrival time
// 3. - 0 = 3
// 12 - 1 = 11
// 18 - 2 = 16

// CASE: FIRST-COME FIRST-SERVED
// 0   1   2   3   4   5   6   7   8   9   10   11   12   13   14   15   16   17   18   19   20   21   22
// s1----------s2------------------------------------s3-----------------------------X        s4----X
// c1----------c1* -> 3
//     c2--------------------------------------------c2* -> 11
//         c3----------------------------------------------------------------------c3* ->
//                                                                                           c4----c4*

// [(0, 3), (1, 9), (2, 6)]

// CASE: OPTIMIZED
// 0   1   2   3   4   5   6   7   8   9   10   11   12   13   14   15   16   17   18   19   20   21   22
// s1----------s3----------------------s2------------------------------------------X
// c1----------c1* (3)
//     c2--------------------------------------------------------------------------c2* (17)
//         c3------c3--------------------c3* (7)

//c1 -> *, c2, c3 -> *
//

// isPrepearing = false
// qw = [], preparing = []
// i = 0
// qw = [c1(3)], qp = []
// qw = [], prepearing = [c1(3)], isPrepearing = true,
// i = 1
// isPreparing = arrivalTime < sandwitchTime
// qw = []
//

package main

import "fmt"

func main() {
	fmt.Println(sandwichRestaurant([][]int{{0, 2}, {2, 4}, {6, 3}}))
}

// [(0, 3), (1, 9), (2, 6)]
func sandwichRestaurant(input [][]int) float64 {
	acumWaitings := 0
	timeWillBeAvailable := 0
	extraWaitingTime := 0

	for i, cusInfo := range input {
		arrivalTime := cusInfo[0]
		sandwichTime := cusInfo[1]

		if arrivalTime < timeWillBeAvailable {
			extraWaitingTime = timeWillBeAvailable - arrivalTime
			timeWillBeAvailable = arrivalTime + extraWaitingTime + sandwichTime
			acumWaitings += extraWaitingTime + sandwichTime
			fmt.Printf("acum: customer %d: %d\n", i, extraWaitingTime+sandwichTime)
		} else {
			timeWillBeAvailable = sandwichTime + arrivalTime
			acumWaitings += sandwichTime
			fmt.Printf("acum: customer %d: %d\n", i, sandwichTime)
		}

	}

	return float64(acumWaitings) / float64(len(input))
}
