package models

import (
	"crypto/sha1"
	"database/sql"
	"fmt"
	"log"
	"todo/config"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB

var err error

const (
	tableNameUser    = "users"
	tableNameMeal    = "meals"
	tableNameSession = "sessions"
)

func init() {
	Db, err = sql.Open(config.Config.SQLDriver, config.Config.DbName)
	if err != nil {
		log.Fatalln(err)
	}

	cmdU := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
id INTEGER PRIMARY KEY AUTOINCREMENT,
uuid STRING NOT NULL UNIQUE,
name STRING,
email STRING,
password STRING,
created_at DATETIME)`, tableNameUser)

	Db.Exec(cmdU)

	cmdM := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		meal_date STRING,
		meal_kind STRING,
		meal_menu STRING,
		meal_place STRING,
		meal_shop_name STRING,
		meal_shop_price INTEGER,
		created_at DATETIME)`, tableNameMeal)

	Db.Exec(cmdM)

	cmdS := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		uuid STRING NOT NULL UNIQUE,
		email STRING,
		user_id INTEGER,
		created_at DATETIME)`, tableNameSession)

	Db.Exec(cmdS)

}

func createUUID() (uuidobj uuid.UUID) {
	uuidobj, _ = uuid.NewUUID()
	return uuidobj
}

func Encrypt(plaintext string) (crypttext string) {
	crypttext = fmt.Sprintf("%s", sha1.Sum([]byte(plaintext)))
	return crypttext
}
