package model

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"time"
)

type Article struct {
	Id    string `json:"id"`
	Title string `json:"title"`
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

func (a ArticleModel) GetAll(title string, genre string, filters Filters) ([]*Article, Metadata, error) {
	query := fmt.Sprintf(`
		SELECT COUNT(*) OVER(), id, title, genre, body, created_at, updated_at 
		FROM articles
		WHERE  (to_tsvector('simple', title) @@ plainto_tsquery('simple', $1) OR $1 = '')
		AND (to_tsvector('simple', genre) @@ plainto_tsquery('simple', $2) OR $2 = '')
		ORDER BY %s %s limit $3 offset $4`, filters.sortColumn(), filters.sortDirection())

	// Create a context with a 3 second timeout
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := a.DB.QueryContext(ctx, query, title, genre, filters.limit(), filters.offset())
	if err != nil {
		return nil, Metadata{}, err
	}

	// Importantly, defer a call to rows.Close() to ensure that the resultset is closed
	// before GetAll() returns.
	defer rows.Close()

	totalRecords := 0
	articles := []*Article{}

	for rows.Next() {
		var article Article

		err := rows.Scan(
			&totalRecords,
			&article.Id,
			&article.Title,
			&article.Genre,
			&article.Body,
			&article.CreatedAt,
			&article.UpdatedAt,
		)
		if err != nil {
			return nil, Metadata{}, err
		}
		articles = append(articles, &article)
	}
	// When the rows.Next() loop has finished, call rows.Err() to retrieve any error
	// that was encountered during the iteration.
	if err = rows.Err(); err != nil {
		return nil, Metadata{}, err
	}

	metadata := calculateMetadata(totalRecords, filters.Page, filters.Page_size)

	return articles, metadata, nil
}

func (a ArticleModel) Insert(article *Article) error {
	query := `
		INSERT INTO articles(title, genre, body)
		VALUES ($1, $2, $3)
		RETURNING id, created_at, updated_at
	`
	args := []interface{}{article.Title, article.Genre, article.Body}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return a.DB.QueryRowContext(ctx, query, args...).Scan(&article.Id, &article.CreatedAt, &article.UpdatedAt)
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

func (a *ArticleModel) Update(article *Article) error {
	query := `
		UPDATE articles
		SET title = $1, genre = $2, body = $3
		WHERE id = $4
		RETURNING updated_at 	
	`
	// FIXME: article.UpdatedAt is not working properly
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
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := a.DB.ExecContext(ctx, query, id)

	return err
}
