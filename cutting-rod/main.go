package main

import (
	"fmt"
	"math"
)

func main() {
	var p = []int{1, 5, 8, 9, 10, 17, 17, 20, 24, 30}
	max := cut(p, 4)

	fmt.Println(max)
}

func cut(p []int, n int) int {
	if n == 0 {
		return 0
	}

	maxValue := int(math.Inf(-1))

	for i := 0; i < n; i++ {
		maxCutI := p[i] + cut(p, n-i-1)
		if maxCutI > maxValue {
			maxValue = maxCutI
		}
	}

	return maxValue
}
