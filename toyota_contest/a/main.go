package main

import (
	"fmt"
	"strconv"
	"strings"
)

// https://atcoder.jp/contests/abc306/tasks/abc306_a
func main() {
	var N, S string
	fmt.Scan(&N, &S)
	n, _ := strconv.Atoi(N)
	if 1 <= n && n <= 50 {
		result := strings.Builder{}
		for _, s := range S {
			result.WriteRune(s)
			result.WriteRune(s)
		}
		fmt.Println(result.String())
	}
}
