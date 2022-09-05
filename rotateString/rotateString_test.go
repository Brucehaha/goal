package rotateString

import (
	"strings"
	"testing"
)

func rotateString(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}
	if len(s) == 0 && len(goal) == 0 {
		return true
	}
	for i := range s {
		substrIn := strings.Index(goal, s[:i+1])
		if substrIn == -1 {
			return false
		}
		if substrIn+i+1 == len(s) {
			return true
		}

	}
	return false
}

func TestRotate(t *testing.T) {
	s, goal := "abcde", "cdeab"
	r := rotateString(s, goal)
	if !r {
		t.Error(" wrong")
	}
}
