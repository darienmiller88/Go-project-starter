package repositories

import (
	"GoProjectStarter/Backend/models"
	"GoProjectStarter/Backend/constants"
	"GoProjectStarter/utils"
	"net/http"

	"github.com/jmoiron/sqlx"
)

type UserRepository interface {
	AddUserDB(name string) models.Result[models.User]
	DeleteUserDB(id int)   models.Result[bool]
	UpdateUserDB()         models.Result[bool]
	GetUsersDB()           models.Result[[]models.User]
}

//Feel free to change this to mongoDB
type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

// AddUser implements [UserRepository].
func (u *userRepository) AddUserDB(name string) models.Result[models.User]{
	id  := 0
	err := u.db.QueryRow(constants.AddUser, name).Scan(&id)

	if err != nil{
		return utils.GetResult(err, http.StatusInternalServerError, models.User{})	
	}
		
	return utils.GetResult(nil, http.StatusOK, models.User{
		ID: id,
		Name: name,
	})
}

// DeleteUser implements [UserRepository].
func (u *userRepository) DeleteUserDB(id int) models.Result[bool]{
	_, err := u.db.Exec(constants.DeleteUser, id)

	if err != nil{
		return utils.GetResult(err, http.StatusInternalServerError, false)	
	}

	return utils.GetResult(nil, http.StatusOK, true)
}

// GetUser implements [UserRepository].
func (u *userRepository) GetUsersDB() models.Result[[]models.User]{
	users := []models.User{}

	if err := u.db.Select(&users, constants.GetUsers); err != nil {
		return utils.GetResult(err, int(http.StatusInternalServerError), []models.User{})
	}

	return utils.GetResult(nil, http.StatusOK, users)
}

// UpdateUser implements [UserRepository].
func (u *userRepository) UpdateUserDB() models.Result[bool]{
	return utils.GetResult(nil, http.StatusOK, true)
}

