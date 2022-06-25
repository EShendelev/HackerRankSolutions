package main

import "fmt"

func breakingRecords(scores []int32) []int32 {
	// Write your code here
	minCount, maxCount := 0, 0
	min, max := scores[0], scores[0]
	for i := 1; i < len(scores); i++ {
		if scores[i] < min {
			min = scores[i]
			minCount++
		} else if scores[i] > max {
			max = scores[i]
			maxCount++
		}
	}
	return []int32{int32(maxCount), int32(minCount)}
}

func main() {
	fmt.Println(breakingRecords([]int32{10, 5, 20, 20, 4, 5, 2, 25, 1}))
	//https://www.hackerrank.com/challenges/breaking-best-and-worst-records/
}
