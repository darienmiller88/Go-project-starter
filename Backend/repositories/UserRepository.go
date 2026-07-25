package repositories

import (
	"GoProjectStarter/Backend/models"
	"GoProjectStarter/utils"
	"net/http"

	"github.com/jmoiron/sqlx"
)

type UserRepository interface {
	AddUser()    models.Result[models.User]
	DeleteUser() models.Result[bool]
	UpdateUser() models.Result[bool]
	GetUser()    models.Result[models.User]
}

//Feel free to change this to mongoDB
type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	return userRepository{
		db: db,
	}
}

// AddUser implements [UserRepository].
func (u userRepository) AddUser() models.Result[models.User]{
	return utils.GetResult(nil, http.StatusOK, models.User{})
}

// DeleteUser implements [UserRepository].
func (u userRepository) DeleteUser() models.Result[bool]{
	return utils.GetResult(nil, http.StatusOK, true)
}

// GetUser implements [UserRepository].
func (u userRepository) GetUser() models.Result[models.User]{
	return utils.GetResult(nil, http.StatusOK, models.User{})
}

// UpdateUser implements [UserRepository].
func (u userRepository) UpdateUser() models.Result[bool]{
	return utils.GetResult(nil, http.StatusOK, true)
}

