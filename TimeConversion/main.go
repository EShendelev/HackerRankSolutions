package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'timeConversion' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */

func timeConversion(s string) string {

	splStr := strings.Split(s[:len(s)-2], ":")

	h, err := strconv.Atoi(splStr[0])
	if err != nil {
		panic("Its not a number")
	}
	m, err := strconv.Atoi(splStr[1])
	if err != nil {
		panic("Its not a number")
	}
	sec, err := strconv.Atoi(splStr[2])
	if err != nil {
		panic("Its not a number")
	}
	if strings.Contains(s, "PM") && h < 12 {
		h += 12
	} else if strings.Contains(s, "AM") && h > 11 {
		h -= 12
	}

	res := fmt.Sprintf("%02d:%02d:%02d", h, m, sec)
	return res

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	s := readLine(reader)

	result := timeConversion(s)

	fmt.Fprintf(writer, "%s\n", result)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
