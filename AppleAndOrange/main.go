package main

import "fmt"

func countApplesAndOranges(s int32, t int32, a int32, b int32, apples []int32, oranges []int32) {
	applesCount, orangesCount := 0, 0
	for i := 0; i < len(apples); i++ {
		if c := a + apples[i]; s <= c && c <= t {
			applesCount++
		}
	}

	for i := 0; i < len(oranges); i++ {
		if c := b + oranges[i]; t >= c && c >= s {
			orangesCount++
		}
	}
	fmt.Printf("%d\n%d\n", applesCount, orangesCount)
}

func main() {

}
