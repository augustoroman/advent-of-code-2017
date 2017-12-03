package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func toi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}

func minmax(s string) (min, max int) {
	nums := strings.Fields(s)
	if len(nums) == 0 {
		return 0, 0
	}

	min = toi(nums[0])
	max = min
	for _, num := range nums {
		n := toi(num)
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}
	return min, max
}

func sum(r io.Reader) int {
	s := bufio.NewScanner(r)
	sum := 0
	for s.Scan() {
		min, max := minmax(s.Text())
		sum += max - min
	}
	if s.Err() != nil {
		panic(s.Err())
	}
	return sum
}

func main() {
	fmt.Println(sum(os.Stdin))
}
