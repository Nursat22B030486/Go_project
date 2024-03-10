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
	// AuthorId  uint `json:"author_id"`
	Genre     string `json:"genre"`
	Body      string `json:"body"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type ArticleModel struct {
	DB       *sql.DB
	InfoLog  *log.Logger
	ErrorLog *log.Logger
}

func (a ArticleModel) Insert(article *Article) {
	query := `
		INSERT INTO articles(title, genre, body)
		VALUES ($1, $2, $3)
	`
	args := []interface{}{article.Title, article.Genre, article.Body}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	a.DB.QueryRowContext(ctx, query, args...)
}

func (a ArticleModel) Get(id int) (*Article, error) {
	query := `
		SELECT id, title, genre, body, created_at, updated_at 
		FROM articles
		where id = $1
	`
	var article Article
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	row := a.DB.QueryRowContext(ctx, query, id)
	err := row.Scan(&article.Id, &article.Title, &article.Genre, &article.Body, &article.CreatedAt, &article.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return &article, nil
}

func (a ArticleModel) Update(article Article) error {
	query := `
		UPDATE articles
		SET title = $1, genre = $2, body = $3,
		WHERE id = $4
		RETURNING updated_at 	
	`

	args := []interface{}{article.Title, article.Genre, article.Body, article.Id}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return a.DB.QueryRowContext(ctx, query, args...).Scan(&article.UpdatedAt)
}

func (a ArticleModel) Delete(id int) error {
	query := `
		DELETE FROM articles
		WHERE id = $1	
	`
	ctx, cancel :=context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := a.DB.ExecContext(ctx, query, id)

	return err
}