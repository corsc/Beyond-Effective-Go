package _3_testability

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/corsc/Beyond-Effective-Go/Chapter04/01_software_design_principles/05_accept_interface_return_struct/03_testability/internal/repo"
	"github.com/corsc/Beyond-Effective-Go/Chapter04/01_software_design_principles/05_accept_interface_return_struct/03_testability/internal/user"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateUser_happyPath(t *testing.T) {
	// mock the database
	db, mockDB, err := sqlmock.New()
	require.NoError(t, err)

	mockDB.ExpectExec("INSERT INTO user").WillReturnResult(sqlmock.NewResult(1, 1))

	// build the repository
	repository := &repo.UserDAO{Database: db}

	// build inputs
	testUser := &user.User{ID: 1, Name: "Amy", Email: "amy@example.com"}

	resultErr := repository.Save(testUser)
	assert.NoError(t, resultErr)

	assert.NoError(t, mockDB.ExpectationsWereMet())
}
