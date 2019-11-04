package main

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type Group struct {
	GroupA string `json:"group_a"`
	GroupB string `json:"group_b"`
}

var groups []Group
var students [][]int

func main() {
	students = [][]int{
		{3},
		{2},
		{1, 4},
		{0, 4, 5},
		{2, 3},
		{3},
	}
	comb := combination(len(students) - 1)

	buildGroups(comb, len(students)-1)

	g, _ := json.MarshalIndent(groups, "", "\t")
	fmt.Println("groups == ", string(g))
	fmt.Println("count == ", len(groups))
}

func applyRule(group Group) bool {
	for i, v := range students {
		isGood := checkEnemy(i, v, group)
		if !isGood {
			return false
		}
	}

	return true
}

func checkEnemy(student int, enemies []int, group Group) bool {
	gA := stringToArray(group.GroupA)
	gB := stringToArray(group.GroupB)

	if inArray(student, gA) {
		for _, e := range enemies {
			if inArray(e, gA) {
				return false
			}
		}
	}

	if inArray(student, gB) {
		for _, e := range enemies {
			if inArray(e, gB) {
				return false
			}
		}
	}

	return true
}

func combination(n int) [][]int {
	if n == 1 {
		return [][]int{
			{0},
			{1},
			{0, 1},
		}
	}

	pre := combination(n - 1)

	var result = make([][]int, len(pre))
	copy(result, pre)

	for _, v := range pre {
		dup := make([]int, len(v))
		copy(dup, v)

		result = append(result, append(dup, n))
	}

	result = append(result, []int{n})

	return result
}

func buildGroups(comb [][]int, n int) {
	for _, gA := range comb {
		gB := buildOtherGroup(gA, n)

		if len(gB) != 0 {
			group := Group{
				GroupA: arrayToString(gA),
				GroupB: arrayToString(gB),
			}

			isGood := applyRule(group)
			if isGood {
				groups = append(groups, group)
			}
		}
	}

}

func arrayToString(array []int) string {
	var result string

	for _, v := range array {
		result += strconv.Itoa(v)
	}

	return result
}

func stringToArray(s string) []int {
	var result []int

	for _, v := range s {
		str := string(v)
		element, err := strconv.Atoi(str)
		if err != nil {
			panic(err)
		}
		result = append(result, element)
	}

	return result
}

func buildOtherGroup(gA []int, n int) []int {
	var gB []int
	for i := 0; i <= n; i++ {
		if !inArray(i, gA) {
			gB = append(gB, i)
		}
	}

	return gB
}

func inArray(e int, array []int) bool {
	for _, v := range array {
		if v == e {
			return true
		}
	}

	return false
}

// A teacher must divide a class of students into two teams to play dodgeball. Unfortunately, not all the kids get along, and several refuse to be put on the same team as that of their enemies.

// Given an adjacency list of students and their enemies, write an algorithm that finds a satisfactory pair of teams, or returns False if none exists.

// For example, given the following enemy graph you should return the teams {0, 1, 4, 5} and {2, 3}.

// students = {
//     0: [3],
//     1: [2],
//     2: [1, 4],
//     3: [0, 4, 5],
//     4: [2, 3],
//     5: [3]
// }
// On the other hand, given the input below, you should return False.

// students = {
//     0: [3],
//     1: [2],
//     2: [1, 3, 4],
//     3: [0, 2, 4, 5],
//     4: [2, 3],
//     5: [3]
// }
