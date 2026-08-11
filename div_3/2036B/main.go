package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func solution(n, k int, b, c []int) int {
	hashMap := make(map[int]int)
	for i := 0; i < len(b); i++ {
		hashMap[b[i]] += c[i]
	}

	values := make([]int, 0, len(hashMap))
	for _, sum := range hashMap {
		values = append(values, sum)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(values)))

	m := len(values)
	limit := min(n, m)

	ans := 0
	for i := 0; i < limit; i++ {
		ans += values[i]
	}

	return ans
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

	t := readInt()
	for i := 0; i < t; i++ {
		n := readInt()
		k := readInt()

		b := make([]int, k)
		c := make([]int, k)
		for j := 0; j < k; j++ {
			b[j] = readInt()
			c[j] = readInt()
		}

		ans := solution(n, k, b, c)
		_, err := fmt.Fprintln(writer, ans)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
	}
}
