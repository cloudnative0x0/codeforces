package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solution(a, b []int) int {
	count := 0

	for i := 0; i < len(a); i++ {
		if i >= len(b) {
			break
		}

		if a[i] > b[i] {
			count++

			b = append(b[:i], b[i+1:]...)
			i--
		} else {
			continue
		}
	}

	return count
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
		b := make([]int, n)

		for j := 0; j < n; j++ {
			a[j] = readInt()
		}

		for f := 0; f < n; f++ {
			b[f] = readInt()
		}

		ans := solution(a, b)
		_, err := fmt.Fprintln(writer, ans)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
	}
}
