package main

import (
	"fmt"
	"math"
)

// https://atcoder.jp/contests/abc305/tasks/abc305_a
func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(math.Round(float64(n)/5) * 5)
}
