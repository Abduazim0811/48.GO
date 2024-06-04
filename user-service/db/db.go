package db

import (
	"social-app/user-service/models"
)

func (s *Storage) CreateUser(user models.User) (models.User, error) {
	var insertUser models.User

	query := s.db.QueryRow("INSERT INTO users(username, email, password) VALUES($1, $2, $3) RETURNING id, username, email",
		user.Username, user.Email, user.Password)
	if err := query.Scan(&insertUser.ID, &insertUser.Username, &insertUser.Email); err != nil {
		return models.User{}, err
	}
	return insertUser, nil
}


// func (s *Storage) GetUser(user models.User) (models.User, error){
// 	var getuser models.User


// }
