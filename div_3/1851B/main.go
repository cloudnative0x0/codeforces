package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
)

func solution(a []int) string {
	copyOfA := make([]int, len(a))
	copy(copyOfA, a)

	slices.Sort(a)

	for i := 0; i < len(a); i++ {
		if (a[i] % 2) != (copyOfA[i] % 2) {
			return "no"
		}
	}

	return "yes"
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
