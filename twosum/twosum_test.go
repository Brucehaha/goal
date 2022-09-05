package twosum

import (
	"fmt"
	"testing"
)

func twoSum(nums []int, target int) []int {
	numsDict := make(map[int]int)
	for i, v := range nums {
		pair_value := target - v
		p, ok := numsDict[pair_value]
		if ok && p != i {
			return []int{i, p}
		}
		numsDict[v] = i

	}
	return []int{}
}

type twoSum3 struct {
	data   map[int]int
	target int
}

func (s *twoSum3) add(num int) {
	if _, ok := s.data[num]; !ok {
		s.data[num] = 1
	} else {
		s.data[num] += 1
	}
}
func (s *twoSum3) find(total int) bool {
	for k, v := range s.data {
		res := total - k
		_, ok := s.data[res]
		if ok {
			if (res == k && v > 1) || (res != k) {
				return true
			}
		}

	}
	return false
}
func TestTwoSum(t *testing.T) {
	a := []int{2, 7, 11}
	result := twoSum(a, 9)
	fmt.Println(result)

}
func TestTwoSum3(t *testing.T) {

	t3 := twoSum3{make(
		map[int]int),
		9,
	}
	t3.add(1)
	t3.add(3)
	t3.add(7)
	ok := t3.find(4)
	if !ok {
		t.Error("not found: 4")
	}
	ok = t3.find(3)
	if !ok {
		t.Error("not found: 3")
	}

}
