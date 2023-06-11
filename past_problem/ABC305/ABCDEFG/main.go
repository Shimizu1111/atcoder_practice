package main

import (
	"fmt"
	"math"
)

// https://atcoder.jp/contests/abc305/tasks/abc305_b
func main() {
	var p, q string
	fmt.Scan(&p, &q)
	resultMap := map[string]int{
		"A": 0,
		"B": 3,
		"C": 4,
		"D": 8,
		"E": 9,
		"F": 14,
		"G": 23,
	}
	fmt.Println(math.Abs(float64(resultMap[p] - resultMap[q])))
}
