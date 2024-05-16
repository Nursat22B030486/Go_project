package model

import (
	"database/sql"
	"log"
	"os"
)

type Models struct {
	Users       UserModel
	Tokens      TokenModel
	Permissions PermissionModel
	Articles    ArticleModel
	Comments    CommentModel
}

func NewModule(db *sql.DB) Models {
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	return Models{
		Users: UserModel{
			DB: db,
		},
		Tokens: TokenModel{
			DB: db,
		},
		Permissions: PermissionModel{
			DB: db,
		},
		Articles: ArticleModel{
			DB:       db,
			InfoLog:  infoLog,
			ErrorLog: errorLog,
		},
		Comments: CommentModel{
			DB:      db,
			InfoLog: infoLog,
			Error:   errorLog,
		},
	}
}
