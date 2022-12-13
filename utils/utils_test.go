package utils

import (
	"testing"
)

func TestJoinListString(t *testing.T) {
	Is(t, JoinListString([]int{1,2,3}, ", "), "1, 2, 3")
}
