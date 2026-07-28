package repositories

import (
	"GoProjectStarter/Backend/models"
	"GoProjectStarter/utils"
	"database/sql"
	"errors"
	"net/http"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func TestAddUserDB(t *testing.T) {
	const testQuery = `INSERT INTO users` // Or exact match/regexp matching constants.AddUser

	t.Run("Happy Path - User successfully created", func(t *testing.T) {
		mock, repo := utils.SetupRepo(t, func(db *sqlx.DB) UserRepository {
			return NewUserRepository(db)
		})

		userName := "Darien"
		expectedID := 1

		// 1. Set mock expectation for QueryRow scan returning the inserted ID
		rows := sqlmock.NewRows([]string{"id"}).AddRow(expectedID)
		mock.ExpectQuery(testQuery).WithArgs(userName).WillReturnRows(rows)

		// 2. Execute method under test
		result := repo.AddUserDB(userName)

		// 3. Assertions
		assert.NoError(t, result.Err)
		assert.Equal(t, http.StatusOK, result.StatusCode)
		assert.Equal(t, expectedID, result.ResultData.ID)
		assert.Equal(t, userName, result.ResultData.Name)

		// Ensure all mock expectations were met
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("Unhappy Path - Database error on insert", func(t *testing.T) {
		mock, repo := utils.SetupRepo(t, func(db *sqlx.DB) UserRepository {
			return NewUserRepository(db)
		})

		userName := "Miller"
		dbError := errors.New("connection failed or unique constraint violated")

		// 1. Expect query to fail with database error
		mock.ExpectQuery(testQuery).WithArgs(userName).WillReturnError(dbError)

		// 2. Execute method under test
		result := repo.AddUserDB(userName)

		// 3. Assertions
		assert.Error(t, result.Err)
		assert.Equal(t, dbError, result.Err)
		assert.Equal(t, http.StatusInternalServerError, result.StatusCode)
		assert.Equal(t, 0, result.ResultData.ID)
		assert.Empty(t, result.ResultData.Name)

		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("Unhappy Path - QueryRow scan fails (e.g. no rows returned)", func(t *testing.T) {
		mock, repo := utils.SetupRepo(t, func(db *sqlx.DB) UserRepository {
			return NewUserRepository(db)
		})

		userName := "Vicky"

		// 1. Simulate sql.ErrNoRows error during Scan
		mock.ExpectQuery(testQuery).WithArgs(userName).WillReturnError(sql.ErrNoRows)

		// 2. Execute method under test
		result := repo.AddUserDB(userName)

		// 3. Assertions
		assert.ErrorIs(t, result.Err, sql.ErrNoRows)
		assert.Equal(t, http.StatusInternalServerError, result.StatusCode)
		assert.Equal(t, models.User{}, result.ResultData)

		assert.NoError(t, mock.ExpectationsWereMet())
	})
}