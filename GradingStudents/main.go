package main

import "fmt"

func gradingStudents(grades []int32) []int32 {
	// Write your code here
	res := make([]int32, 0)
	for _, v := range grades {
		res = append(res, finalGrading(v))
	}
	return res
}

func finalGrading(a int32) int32 {
	nextMult := a - a%5 + 5
	if a < 38 {
		return a
	}

	if nextMult-a < 3 {
		return nextMult
	} else if nextMult-a >= 3 {
		return a
	}
	return 0
}

func main() {
	grades := []int32{73, 67, 38, 37, 33}
	fmt.Println(gradingStudents(grades))
}
