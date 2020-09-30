package _1_greetings

func buildGreeting(name string) func() string {
	greeting := "Hello " + name + "."

	return func() string {
		out := greeting + " Nice to meet you!"
		return out
	}
}
