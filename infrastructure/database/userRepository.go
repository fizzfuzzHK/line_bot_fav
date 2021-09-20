package database

import (
	domain "github.com/fizzfuzzHK/line_bot_fav/domain"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/jmoiron/sqlx" // mysql
)

type IUserRepository interface {
	AddUser(UserID string) string
}

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (repo *UserRepository) AddUser(UserID string) string {
	user := domain.User{}
	err := repo.db.Get(&user, "SELECT * FROM users WHERE user_id = ? LIMIT 1", UserID)
	if err != nil {
		query := "Insert INTO users (user_id) VALUES (?)"
		repo.db.Queryx(query, UserID)
		return UserID
	} else {
		return ""
	}
}
