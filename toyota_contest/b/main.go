package main

import (
	"fmt"
	"math"
	"strconv"
)

// https://atcoder.jp/contests/abc306/tasks/abc306_b
// このコードはNG
func main() {
	var inputs []int
	count := -1
	for {
		if count == 3 {
			break
		}
		var input string
		fmt.Scan(&input)
		t, _ := strconv.Atoi(input)
		inputs = append(inputs, t)
		count++
	}

	result := 0
	for i, v := range inputs {
		if i == 0 {
			result += 1
		}
		if v == 1 && i != 0 {
			result += int(math.Pow(2, float64(i)))
		}
	}
	rts := strconv.Itoa(result)

	fmt.Println(rts)
}
