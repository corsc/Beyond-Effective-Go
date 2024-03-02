package _1_minimalist_modular

func Example_recordUserLogin() {
	// Output: UserID 123 logged in
	recordUserLogin(fred)
}

func Example_recordUserLoginByID() {
	// Output: UserID 123 logged in
	recordUserLoginByID(fred)
}

var fred = &User{
	ID:    123,
	Name:  "Fred",
	Email: "fred@example.com",
}
