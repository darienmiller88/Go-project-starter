package constants

const (
	
	GetUsers string = "SELECT * FROM users"

	DeleteUser string = "DELETE FROM users WHERE id = $1"

	AddUser string = "INSERT INTO users (name) VALUES ($1)"

	UpdateUser string = ""
)