package database

import (
	"fmt"
	"os"
	"database/sql"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	Db *sql.DB
	er error
)

func Init() {
	MYSQL_USER := os.Getenv("MYSQL_USER")
	MYSQL_PASSWORD := os.Getenv("MYSQL_PASSWORD")
	PROTOCOL := "tcp(" + os.Getenv("MYSQL_DATABASE_HOST") + ":" + "3306)"
	MYSQL_DATABASE := os.Getenv("MYSQL_DATABASE")

	CONNECT := fmt.Sprintf(
		"%s:%s@%s/%s?charset=utf8mb4&parseTime=True&loc=Local",
		MYSQL_USER,
		MYSQL_PASSWORD,
		PROTOCOL,
		MYSQL_DATABASE,
	)
	fmt.Println(CONNECT)
	connectDb, er := gorm.Open(mysql.Open(CONNECT), &gorm.Config{})
	if er != nil {
		panic(er)
	}
	Db, er = connectDb.DB()
	fmt.Println("変数の中身: ", Db)
}
