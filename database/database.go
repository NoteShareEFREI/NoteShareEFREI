package database

import (
	"database/sql"
	"fmt"
	"os"
)

// Db global database connection, should be used everywhere we need to query data
var Db *sql.DB

func ConnectToDBUnix() (db *sql.DB, err error) {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbSocketPath := os.Getenv("DB_SOCKET_PATH")
	dbName := os.Getenv("DB_NAME")

	return sql.Open("mysql", fmt.Sprintf("%s:%s@unix(%s)/%s?parseTime=true", dbUser, dbPassword, dbSocketPath, dbName))
}

func ConnectToDBUnixLegacy(dbUser string, dbPassword string, dbSocketPath string, dbName string) (db *sql.DB, err error) {
	return sql.Open("mysql", fmt.Sprintf("%s:%s@unix(%s)/%s?parseTime=true", dbUser, dbPassword, dbSocketPath, dbName))
}

func ConnectToDBTcp() (db *sql.DB, err error) {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	return sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPassword, dbHost, dbPort, dbName))
}
