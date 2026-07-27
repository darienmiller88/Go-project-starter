package repositories

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/require"
)

func setupUserRepo(t *testing.T) (sqlmock.Sqlmock, UserRepository) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)

	sqlxDB := sqlx.NewDb(db, "postgres")
	repo := NewUserRepository(sqlxDB)

	t.Cleanup(func() {
		db.Close()
	})

	return mock, repo
}