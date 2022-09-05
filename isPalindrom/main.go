package main

import (
	"fmt"
	"math"
	"strconv"
	"unicode"
)

func isPalindrom(s string) bool {
	i, j := 0, len(s)-1
	for i <= j && i < len(s)-1 && j >= 0 {
		si := int(s[i])
		if (si < 64 && si > 57) || si < 48 || (90 < si && si < 97) || si > 122 {
			i++
			continue
		}
		sj := int(s[j])
		if (sj < 65 && sj > 57) || sj < 48 || (90 < sj && sj < 97) || sj > 122 {
			j--
			continue
		}
		if si != sj {
			if math.Abs(float64(si)-float64(sj)) == float64(32) {
				if si < 65 || (si > 90 && si < 97) || sj < 65 || (sj > 90 && sj < 97) || si > 122 || sj > 122 {
					return false
				}
			} else {
				return false
			}
		}
		i++
		j--

	}
	return true
}
func isPalindrom2(s string) bool {
	i, j := 0, len(s)-1
	for i < j {
		si := rune(s[i])
		if !unicode.IsLetter(si) && !unicode.IsNumber(si) {
			i++
		}

		sj := rune(s[j])
		if !unicode.IsLetter(sj) && !unicode.IsNumber(sj) {
			j--
		}
		if unicode.ToLower(sj) != unicode.ToLower(si) {
			return false
		}
		i++
		j--

	}
	return true
}
func main() {
	s := "aA"
	fmt.Println(isPalindrom2(s))
	for _, a := range []byte("abc你好") {
		fmt.Printf("%T, %v \n", a, a)
	}

	bs := []byte(strconv.Itoa(31415926))
	fmt.Println(bs)
	numsDict := make(map[int]int)

	fmt.Printf("%V", numsDict)
}
