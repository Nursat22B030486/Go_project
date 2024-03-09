package model

import (
	"database/sql"
	"log"
)

type UserModel struct {
	DB       *sql.DB
	InfoLog  *log.Logger
	ErrorLog *log.Logger
}

func NewUserModel(db *sql.DB) *UserModel {
	return &UserModel{DB: db}
}

type User struct {
	Id        int    `json:"id"`
	Fullname  string `json:"full_name"`
	Username  string `json:"username"`
	Password  string `json:"passsword"`
	CreatedAt string `json:"created_at"`
}
