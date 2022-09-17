package main

import "fmt"

func main() {
	var w, v []int
	var c, n int
	c = 5
	w = []int{1, 3, 3}
	v = []int{6, 10, 12}
	n = len(v)

	// r := calculate_recursive(w, v, c, n)
	r := dp(w, v, c, n)
	fmt.Println(r)
}

func dp(w, v []int, c, n int) int {
	return 0
}

func calculate_recursive(w, v []int, c, n int) int {
	// fmt.Println(w, v, c, n)
	if n == 0 || c <= 0 { //cap can't be negative because of the next if statement. So it should be c == 0
		return 0
	}

	// if previous item is heavier than the cap, must ignore it!
	if w[n-1] > c {
		return calculate_recursive(w, v, c, n-1)
	}

	exclude := calculate_recursive(w, v, c, n-1)
	include := v[n-1] + calculate_recursive(w, v, c-w[n-1], n-1)

	return max(exclude, include)

	// include := v[n-1] + calculate(w, v, c-w[n-1], n-1)

	// fmt.Println("include = ", include)
	// return include
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
