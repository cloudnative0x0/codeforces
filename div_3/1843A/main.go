package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func solution(a []int) int {
	sort.Ints(a)

	maxVal := 0
	left, right := 0, len(a)-1

	for left < right {
		maxVal = a[right] - a[left]

		left++
		right--
	}

	return maxVal
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

		val := solution(a)

		_, err := fmt.Fprintln(writer, val)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
	}
}
