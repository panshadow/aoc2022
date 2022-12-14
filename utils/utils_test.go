package utils

import (
	"testing"
)

func TestJoinListString(t *testing.T) {
	Is(t, JoinListString([]int{1,2,3}, ", "), "1, 2, 3")
	IsNamed(t, "Bad fload", JoinListString([]float32{1.0,2.1,3.14}, " :: "), "1. :: 2.1 :: 3.14")
}
