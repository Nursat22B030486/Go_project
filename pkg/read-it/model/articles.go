package model

import (
	"context"
	"database/sql"
	"log"
	"time"
)

type Article struct {
	Id        string `json:"id"`
	Title     string `json:"title"`
	AuthorId  string `json:"author_id"`
	Genre     string `json:"genre"`
	Body      string `json:"text"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type ArticleModel struct {
	DB       *sql.DB
	InfoLog  *log.Logger
	ErrorLog *log.Logger
}

func (a ArticleModel) Insert(article *Article) error {
	query := `
		INSERT INTO "Articles"(title, author_id, genre, body)
		VALUES ($1, $2, $3, $4)
		RETURNING id, created_at, update_at
	`
	args := []interface{}{article.Title, article.AuthorId, article.Genre, article.Body}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return a.DB.QueryRowContext(ctx, query, args...).Scan(&article.Id, &article.CreatedAt, &article.UpdatedAt)
}

func (a ArticleModel) Get(id int) (*Article, error) {
	query := `
		SELECT id, title, author_id, genre, body, created_at, updated_at 
		FROM "Articles"
		where id = $1
	`
	var article Article
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	row := a.DB.QueryRowContext(ctx, query, id)
	err := row.Scan(&article.Id, &article.Title, &article.AuthorId, &article.Genre, &article.Body, &article.CreatedAt, &article.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return &article, nil
}

func (a ArticleModel) Update(article Article) error {
	query := `
		UPDATE "Articles"
		SET title = $1, author_id = $2, genre = $3, body = $4
		WHERE id = $5
		RETURNING updated_at 	
	`

	args := []interface{}{article.Title, article.AuthorId, article.Genre, article.Body, article.Id}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return a.DB.QueryRowContext(ctx, query, args...).Scan(&article.UpdatedAt)
}

func (a ArticleModel) Delete(id int) error {
	query := `
		DELETE FROM "Articles"
		WHERE id = $1	
	`
	ctx, cancel :=context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := a.DB.ExecContext(ctx, query, id)

	return err
}