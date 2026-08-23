package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solution(a []int) (int64, int) {
	var maxSum int64
	count := 0

	for i := 0; i < len(a); i++ {
		if a[i] < 0 {
			maxSum += int64(-a[i])
		} else {
			maxSum += int64(a[i])
		}

		if a[i] < 0 {
			count++

			for i+1 < len(a) && a[i+1] <= 0 {
				i++

				if a[i] < 0 {
					maxSum += int64(-a[i])
				}
			}
		}
	}

	return maxSum, count
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1024*1024), 1024*1024)
	scanner.Split(bufio.ScanWords)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer func(writer *bufio.Writer) {
		err := writer.Flush()
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
	}(writer)

	readInt := func() int {
		scanner.Scan()
		x, _ := strconv.Atoi(scanner.Text())

		return x
	}

	q := readInt()
	for i := 0; i < q; i++ {
		n := readInt()
		a := make([]int, n)

		for j := 0; j < n; j++ {
			a[j] = readInt()
		}

		sum, count := solution(a)

		_, err := fmt.Fprintln(writer, sum, count)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
	}
}
