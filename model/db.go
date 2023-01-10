package model

import (
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	sql "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

func InitDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_DATABASE"))

	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	log.Println("db connected!")

	DB.AutoMigrate(&User{}, &Profile{})

	return DB, err
}

func InitDBMock() (*gorm.DB, sqlmock.Sqlmock, error) {
	MockDB, mock, err := sqlmock.New()

	if err != nil {
		return nil, mock, err
	}

	db, err := gorm.Open(mysql.New(
		mysql.Config{
			Conn:                      MockDB,
			DSNConfig:                 &sql.Config{ParseTime: true},
			SkipInitializeWithVersion: true,
		}),
		&gorm.Config{},
	)

	return db, mock, err
}
