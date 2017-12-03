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

func minmax(line string) (min, max int) {
	nums := lineToNums(line)
	if len(nums) == 0 {
		return 0, 0
	}
	min, max = nums[0], nums[0]
	for _, n := range nums {
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}
	return min, max
}

func readLines(r io.Reader) []string {
	var lines []string
	s := bufio.NewScanner(r)
	for s.Scan() {
		line := strings.TrimSpace(s.Text())
		if len(line) > 0 {
			lines = append(lines, line)
		}
	}
	if s.Err() != nil {
		panic(s.Err())
	}
	return lines
}

func sum(lines []string) int {
	sum := 0
	for _, line := range lines {
		min, max := minmax(line)
		sum += max - min
	}
	return sum
}

func lineToNums(s string) []int {
	parts := strings.Fields(s)
	nums := make([]int, len(parts))
	for i, numstr := range parts {
		nums[i] = toi(numstr)
	}
	return nums
}

func firstEvenDivisiblesRes(nums []int) (result int) {
	for i, a := range nums {
		for j := i + 1; j < len(nums); j++ {
			b := nums[j]
			if a > b && a%b == 0 && b != 0 {
				return a / b
			} else if a < b && b%a == 0 && a != 0 {
				return b / a
			}
		}
	}
	panic(fmt.Errorf("no evenly divisible numbers in %v", nums))
}

func sum2(lines []string) int {
	sum := 0
	for _, line := range lines {
		sum += firstEvenDivisiblesRes(lineToNums(line))
	}
	return sum
}

func main() {
	lines := readLines(os.Stdin)
	fmt.Println("sum1:", sum(lines))
	fmt.Println("sum2:", sum2(lines))
}
