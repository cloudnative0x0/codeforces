package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func solution(a []int) string {
	sort.Ints(a)
	left, right := 1, len(a)-1

	blueSum := a[0] + a[1]
	redSum := a[right]

	for left < right {
		if redSum > blueSum {
			return "yes"
		}

		left++
		right--

		blueSum += a[left]
		redSum += a[right]
	}

	return "no"
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

		ans := solution(a)

		_, err := fmt.Fprintln(writer, ans)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
	}
}
