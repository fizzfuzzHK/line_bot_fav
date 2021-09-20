package infrastructure

import (
	"os"

	domain "github.com/fizzfuzzHK/line_bot_fav/domain"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/jmoiron/sqlx" // mysql
	"github.com/sirupsen/logrus"
)

type IUserRepository interface {
	AddUser(UserID string) string
}

type UserRepository struct {
	db *sqlx.DB
}

func NewDbClient(db *sqlx.DB) *UserRepository {
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

// Connect DB接続
func Connect() (db *sqlx.DB, err error) {
	db, err = sqlx.Open("mysql",
		os.Getenv("DB_USERNAME")+":"+os.Getenv("DB_USERPASS")+
			"@tcp("+os.Getenv("DB_HOST")+":"+os.Getenv("DB_PORT")+")/"+
			os.Getenv("DB_NAME")+
			"?charset=utf8mb4&parseTime=True&loc=Local")

	if err != nil {
		logrus.Fatalf("Error connect DB: %v", err)
	}

	return db, err
}
