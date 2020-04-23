// https://www.hackerrank.com/challenges/marcs-cakewalk/problem

package main

import (
	"fmt"
	"math"
	"sort"
)

func marcsCakewalk(calorie []int32) int64 {
	sort.Slice(calorie, func(i, j int) bool {
		return calorie[j] < calorie[i]
	})

	var min = int64(0)

	for i := 0; i < len(calorie); i++ {
		min += int64(math.Pow(2.0, float64(i))) * int64(calorie[i])
	}

	return min
}

func main() {
	fmt.Println(marcsCakewalk([]int32{1, 3, 2}))
	fmt.Println(marcsCakewalk([]int32{5, 10, 7}))
	fmt.Println(marcsCakewalk([]int32{7, 4, 9, 6}))
}
