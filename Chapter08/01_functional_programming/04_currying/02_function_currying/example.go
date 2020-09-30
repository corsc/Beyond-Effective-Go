package _2_function_currying

func Multiply(a int) func(int) int {
	return func(b int) int {
		return a * b
	}
}
