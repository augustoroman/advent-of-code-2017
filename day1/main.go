package main

import (
	"fmt"
	"os"
)

func sumRepeatedDigits(d string) int {
	sum := 0
	if len(d) == 0 {
		return sum
	}
	prev := rune(d[len(d)-1])
	for _, ch := range d {
		if ch == prev && ch >= '0' && ch <= '9' {
			sum += int(ch - '0')
		}
		prev = ch
	}
	return sum
}

func sumOffsetRepeatedDigits(d string) int {
	sum := 0
	N := len(d)
	if N == 0 {
		return sum
	}
	offset := N / 2
	get := func(pos int) rune { return rune(d[(pos+offset)%N]) }
	for pos, ch := range d {
		ch2 := get(pos)
		if ch == ch2 && ch >= '0' && ch <= '9' {
			sum += int(ch - '0')
		}
	}
	return sum
}

func main() {
	for _, str := range os.Args[1:] {
		fmt.Printf("Sum1 = %d\n", sumRepeatedDigits(str))
		fmt.Printf("Sum2 = %d\n", sumOffsetRepeatedDigits(str))
	}
}
