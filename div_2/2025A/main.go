package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solution(s1, s2 string) int {
	k := 0

	for i := 0; i < min(len(s1), len(s2)); i++ {
		if s1[i] == s2[i] {
			k++
		} else {
			break
		}
	}

	if k == 0 {
		return len(s1) + len(s2)
	}

	return len(s1) + len(s2) - k + 1
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
		s1 := readString()
		s2 := readString()

		ans := solution(s1, s2)
		_, err := fmt.Fprintln(writer, ans)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
	}
}
