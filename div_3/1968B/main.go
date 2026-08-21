package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solution(a, b string) int {
	moveA, moveB := 0, 0

	for moveA < len(a) && moveB < len(b) {
		if a[moveA] == b[moveB] {
			moveA++
		}

		moveB++
	}

	return moveA
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
		_ = readInt()

		a := readString()
		b := readString()

		ans := solution(a, b)

		_, err := fmt.Fprintln(writer, ans)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
	}
}
