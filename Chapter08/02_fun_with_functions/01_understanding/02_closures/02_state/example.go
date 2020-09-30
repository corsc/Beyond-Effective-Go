package _2_state

func buildTotaler() func(int) int {
	total := 0
	return func(in int) int {
		total += in
		return total
	}
}
