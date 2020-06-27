package _3_testability

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/corsc/Advanced-Go-Programming/Chapter04/01_software_design_principles/05_accept_interface_return_struct/03_testability/internal/repo"
	"github.com/corsc/Advanced-Go-Programming/Chapter04/01_software_design_principles/05_accept_interface_return_struct/03_testability/internal/user"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateUser_happyPath(t *testing.T) {
	// mock the database
	db, mockDB, err := sqlmock.New()
	require.NoError(t, err)

	mockDB.ExpectExec("INSERT INTO user (id,name,email) VALUES (?,?,?)").
		WillReturnResult(sqlmock.NewResult(1, 1))

	// build the repository
	repository := &repo.UserDAO{Database: db}

	// build inputs
	testUser := &user.User{ID: 1, Name: "Amy", Email: "amy@home.com"}

	resultErr := repository.Save(testUser)
	assert.NoError(t, resultErr)
}
