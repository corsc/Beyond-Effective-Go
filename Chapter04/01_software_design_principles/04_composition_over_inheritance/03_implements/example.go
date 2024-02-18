package _3_implements

// Use the compiler to ensure Duck is a talker
var duckTalker talker = Duck{}

type talker interface {
	Talk() string
}

type Duck struct {
}

func (d Duck) Talk() string {
	return "Quack!"
}
