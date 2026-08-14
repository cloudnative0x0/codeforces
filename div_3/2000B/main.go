package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solution(n int, nums []int) string {
	bus := make([]bool, n+2)

	for i := 0; i < n; i++ {
		seat := nums[i]

		if i > 0 {
			if !bus[seat-1] && !bus[seat+1] {
				return "NO"
			}
		}

		bus[seat] = true
	}

	return "YES"
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
		nums := make([]int, n)

		for j := 0; j < n; j++ {
			nums[j] = readInt()
		}

		ans := solution(n, nums)
		_, err := fmt.Fprintln(writer, ans)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
	}
}
