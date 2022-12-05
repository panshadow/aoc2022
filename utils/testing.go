package utils

import (
	"fmt"
	"testing"
)

func RunTest(input []string, task TaskSolutionFn, expected string) string {
	result := task(input)
	if result != expected {
		return fmt.Sprintf("Actual: %s\nExpected: %s\n", result, expected)
	}
	return ""
}

func Is[T comparable](t *testing.T, actual T, expected T) {
	if actual != expected {
		t.Errorf("FAIL\nActual: %v\nExpected: %v\n", actual, expected)
	}
}
