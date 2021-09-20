package infrastructure

import (
	"os"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/jmoiron/sqlx" // mysql
	"github.com/sirupsen/logrus"
)

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
