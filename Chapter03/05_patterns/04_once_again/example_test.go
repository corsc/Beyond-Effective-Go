package _4_once_again

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

// go test -run=Bench. -bench=. -benchtime=10s ./Chapter03/05_patterns/04_once_again/
// goos: darwin
// goarch: amd64
// pkg: github.com/PacktPublishing/Advanced-Go-Programming/Chapter03/05_patterns/04_once_again
// BenchmarkExample-8   	 2000000	      8922 ns/op
// BenchmarkFixed-8     	 2000000	      8777 ns/op

func BenchmarkExample(b *testing.B) {
	// build and configure the mock DB
	db, mockDB, _ := sqlmock.New()
	mockDB.ExpectExec("INSERT INTO user (id,name,email) VALUES (?,?,?)").
		WillReturnResult(sqlmock.NewResult(1, 1))

	// build inputs
	dao := &UserDAO{}
	testUser := User{ID: 2, Name: "Test", Email: "test@example.com"}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = dao.Save(db, testUser)
	}
}

func BenchmarkFixed(b *testing.B) {
	// build and configure the mock DB
	db, mockDB, _ := sqlmock.New()
	mockDB.ExpectExec("INSERT INTO user (id,name,email) VALUES (?,?,?)").
		WillReturnResult(sqlmock.NewResult(1, 1))

	// build inputs
	dao := &UserDAO{}
	testUser := User{ID: 2, Name: "Test", Email: "test@example.com"}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = dao.SaveFixed(db, testUser)
	}
}
