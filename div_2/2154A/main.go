package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solution(s string, k int) int {
	ans := 0
	last := -k

	for i := 0; i < len(s); i++ {
		if s[i] == '1' {
			if i-last >= k {
				ans++
			}

			last = i
		}
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
	readString := func() string {
		scanner.Scan()
		return scanner.Text()
	}

	q := readInt()
	for i := 0; i < q; i++ {
		_ = readInt()
		k := readInt()

		s := readString()

		ans := solution(s, k)

		_, err := fmt.Fprintln(writer, ans)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
	}
}
