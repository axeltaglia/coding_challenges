package main

import "fmt"

func main() {
	//canCompleteCircuit([]int{3, 1, 1, 1, 5}, []int{2, 2, 2, 2, 1})
	fmt.Println(canCompleteCircuit([]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2}))
	//fmt.Println(canCompleteCircuit([]int{6, 1, 4, 3, 5}, []int{3, 8, 2, 4, 2}))
	//canCompleteCircuit([]int{2, 3, 4}, []int{3, 4, 3})
}

func canCompleteCircuit(gas []int, cost []int) int {
	aux := make([]int, len(gas))

	acum := 0
	for i, _ := range gas {
		aux[i] = gas[i] - cost[i]
		acum += aux[i]
	}

	if acum < 0 {
		return -1
	}
	// we can make sure there is possible to complete the circuit, and we have a unique solution.

	fmt.Println(aux)
	cand := -1
	acum = 0
	for i, n := range aux {
		acum += n
		if cand == -1 && n >= 0 {
			cand = i
			acum = n
		} else if acum < 0 {
			cand = -1
		}
	}

	return cand
}

// si no hay candidato y n > 0 --> cand = i
// si hay candidato y acum < 0 --> cand = -1
//

func canCompleteCircuitPOC(gas []int, cost []int, start int) int {
	acum := gas[start]
	it := 0
	i := start
	fmt.Println(acum)
	for it < len(gas) {
		next := i + 1
		if next == len(gas) {
			next = 0
		}
		fmt.Printf("acum = %d - %d + %d\n", acum, cost[i], gas[next])
		acum = acum - cost[i] + gas[next]
		i = next
		fmt.Println(acum)
		it++
	}

	return 2
}

//  1   2   3   4   5
//  3   4   5   1   2
// -2  -2  -2   3   3

// 5 - 2 + 1 - 3 + 2 - 4 + 3 - 5 + 4 - 1 + 5

// 3     + -2    + -2    +  -2   +  3    + 3

// 3   +   3    +  -2    +  -2   + -2    + 3

// 2   3    4
// 3   4    3
// -1  -1   1

// 3   1   1   1   5
// 2   2   2   2   1
// 1  -1  -1  -1   4             -1  4  1  -1  -1
// 1  0   -1  -2   2

// 2   3   1
// 3   1   2
// -1  2  -1

// canP = -1 (there is no candidate)
// canP = 2 nums[canP] > 0
// acum

// 1,1,5,2,3
// 4,3,1,2,2

// -3  -2  4  0  1
//

// -3  -2  4  0  1  -->  2

// -1  3  -4  2

// -2  -2  -2  3  3  --> 3

// 1 -1 -1 -1 4 --> 4

// si no hay candidato Y n > 0 --> cand = n
// si hay candidato y acum < 0 --> cand = -1
// si hay candidato y acum < 0

// 3  -7  2  -1  3
// 3  -1  2   2  2
// 3  -4 -2  -3  0
