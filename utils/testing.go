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

func IsNamed[T comparable](t *testing.T, title string, actual T, expected T) {
	if actual != expected {
		t.Errorf("FAIL %s\nActual: %v\nExpected: %v\n", title, actual, expected)
	}
}

func Is[T comparable](t *testing.T, actual, expected T) {
	IsNamed(t, "", actual, expected)
}
