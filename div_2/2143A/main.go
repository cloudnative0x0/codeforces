package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solution(n int, p []int) string {
	arr := make([]int, n)
	copy(arr, p)

	for k := 1; k <= n; k++ {
		start := -1
		end := -1
		for i := 0; i < n; i++ {
			if arr[i] > n-k {
				if start == -1 {
					start = i
				}

				end = i
			}
		}

		if start == -1 {
			start = 0
			end = k - 1
		}

		if end-start+1 > k {
			return "NO"
		}

		for end-start+1 < k {
			if end < n-1 {
				end++
			} else if start > 0 {
				start--
			}
		}

		for i := start; i <= end; i++ {
			arr[i]--
		}
	}

	for _, v := range arr {
		if v != 0 {
			return "NO"
		}
	}

	return "YES"
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1024*1024), 1024*1024)
	scanner.Split(bufio.ScanWords)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	readInt := func() int {
		scanner.Scan()
		x, _ := strconv.Atoi(scanner.Text())

		return x
	}

	t := readInt()
	for i := 0; i < t; i++ {
		n := readInt()
		p := make([]int, n)

		for j := 0; j < n; j++ {
			p[j] = readInt()
		}

		ans := solution(n, p)
		fmt.Fprintln(writer, ans)
	}
}
