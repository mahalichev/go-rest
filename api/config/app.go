package config

import (
	"log"
	"server/api/models/sql"
)

type Application struct {
	ErrLog  *log.Logger
	InfoLog *log.Logger
	Users   *sql.UserModel
}
