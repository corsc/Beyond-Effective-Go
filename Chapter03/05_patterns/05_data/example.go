package _5_data

//go:generate protoc -I=. --go_out=. fixed.proto
type User struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	City    string `json:"city"`
	Country string `json:"country"`
}
