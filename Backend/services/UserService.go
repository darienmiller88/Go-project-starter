package services

import (
	"GoProjectStarter/Backend/models"
	"GoProjectStarter/Backend/repositories"
)

type UserService interface {
	AddUser(name string) models.Result[models.User]
	DeleteUser(id int)   models.Result[bool]
	UpdateUser()         models.Result[bool]
	GetUsers()           models.Result[[]models.User]
}

type userService struct {
	repo repositories.UserRepository	
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{
		repo: repo,
	}
}

// AddUser implements [UserService].
func (u *userService) AddUser(name string) models.Result[models.User] {
	return u.repo.AddUserDB(name)
}

// DeleteUser implements [UserService].
func (u *userService) DeleteUser(id int) models.Result[bool] {
	return u.repo.DeleteUserDB(id)
}

// GetUsers implements [UserService].
func (u *userService) GetUsers() models.Result[[]models.User] {
	return u.repo.GetUsersDB()
}

// UpdateUser implements [UserService].
func (u *userService) UpdateUser() models.Result[bool] {
	return u.repo.UpdateUserDB()
}

