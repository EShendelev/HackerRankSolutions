package main

func kangaroo(x1 int32, v1 int32, x2 int32, v2 int32) string {
	if v1 <= v2 {
		return "NO"
	} else if (x1-x2)%(v2-v1) == 0 {
		return "YES"
	}
	return "NO"

}

func main() {
	// https://www.hackerrank.com/challenges/kangaroo
}
