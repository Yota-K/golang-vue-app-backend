package database

import (
	"fmt"
	"os"
)

func Database_init() {
	fmt.Println("DB_HOST:", os.Getenv("MYSQL_DATABASE_HOST"))
	fmt.Println("DB_NAME:", os.Getenv("MYSQL_DATABASE"))
	fmt.Println("DB_USER:", os.Getenv("MYSQL_USER"))
	fmt.Println("DB_PASS:", os.Getenv("MYSQL_PASSWORD"))
}
