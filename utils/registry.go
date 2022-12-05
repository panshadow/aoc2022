package utils

type TaskSolutionFn func([]string) string

type TaskSolution struct {
	Code       TaskSolutionFn
	Input string
}

var (
	TaskRegistry = map[string]TaskSolution{}
)

func RegisterTask(taskID string, code TaskSolutionFn, input string) {
	if input == "" {
		input = taskID
	}
	TaskRegistry[taskID] = TaskSolution{
		Code: code,
		Input: input,
	}
}
