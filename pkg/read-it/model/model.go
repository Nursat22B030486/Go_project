package model

import (
	"database/sql"
	"log"
	"os"
)

type Models struct {
	Articles ArticleModel
	Users UserModel
}

func NewModule(db *sql.DB) Models{
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	return Models {
		Articles: ArticleModel{
			DB: db,
			InfoLog: infoLog,
			ErrorLog: errorLog,
		},
		Users: UserModel{
			DB: db,
			InfoLog: infoLog,
			ErrorLog: errorLog,
		},
	}
}