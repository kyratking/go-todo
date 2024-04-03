package config

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ichtrojan/thoth"
	_ "github.com/joho/godotenv/autoload"
)

func Database() *sql.DB {
	logger, _ := thoth.Init("log")

	user, exist := os.LookupEnv("MYSQL_USER")

	if !exist {
		logger.Log(errors.New("MYSQL_USER not set in .env"))
		log.Fatal("DB_USER not set in .env")
	}

	password, exist := os.LookupEnv("MYSQL_PASSWORD")

	if !exist {
		logger.Log(errors.New("MYSQL_PASSWORD not set in .env"))
		log.Fatal("DB_PASSWORD not set in .env")
	}

	host, exist := os.LookupEnv("MYSQL_HOST")

	if !exist {
		logger.Log(errors.New("MYSQL_HOST not set in .env"))
		log.Fatal("DB_HOST not set in .env")
	}

	credentials := fmt.Sprintf("%s:%s@(%s:3306)/?charset=utf8&parseTime=True", user, password, host)

	database, err := sql.Open("mysql", credentials)

	if err != nil {
		logger.Log(err)
		log.Fatal(err)
	} else {
		fmt.Println("Database Connection Successful")
	}

	_, err = database.Exec("CREATE DATABASE IF NOT EXISTS todo")

	if err != nil {
		logger.Log(err)
		log.Fatal(err)
	} else {
		fmt.Println("Database Created")
	}

	_, err = database.Exec("USE todo")

	if err != nil {
		logger.Log(err)
		log.Fatal(err)
	} else {
		fmt.Println("Database Selected")
	}

	_, err = database.Exec("CREATE TABLE IF NOT EXISTS todos (id INT AUTO_INCREMENT PRIMARY KEY, title VARCHAR(100), description TEXT, created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP)")

	if err != nil {
		logger.Log(err)
		log.Fatal(err)
	} else {
		fmt.Println("Table Created")
	}

	return database
}
