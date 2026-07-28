package utils

import (
	"GoProjectStarter/Backend/models"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/require"
)

// Helper function to allow repos to send result payloads with less text.
func GetResult[T any](err error, statusCode int, payload T) models.Result[T] {
	return models.Result[T]{
		StatusCode: statusCode,
		Err:        err,
		ResultData: payload,
	}
}

// SetupRepo creates a mocked sqlx.DB and initializes any repository type T using a factory function.
func SetupRepo[T any](t *testing.T, factory func(db *sqlx.DB) T) (sqlmock.Sqlmock, T) {
	t.Helper() // Marks this function as a test helper for accurate file/line reporting in failures

	db, mock, err := sqlmock.New()
	require.NoError(t, err)

	sqlxDB := sqlx.NewDb(db, "postgres")
	
	// Instantiate whatever repository T is passed via the factory closure
	repo := factory(sqlxDB)

	t.Cleanup(func() {
		db.Close()
	})

	return mock, repo
}