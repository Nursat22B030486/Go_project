package model

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"time"
)

type Comment struct {
	Id        string `json:"id"`
	ArticleId uint   `json:"article_id"`
	UserId    uint   `json:"user_id"`
	Body      string `json:"body"`
	CreatedAt string `json:"createdAt"`
}

type CommentModel struct {
	DB      *sql.DB
	InfoLog *log.Logger
	Error   *log.Logger
}

func (c CommentModel) GetAll(article_id int, body string, filters Filters) ([]*Comment, Metadata, error) {
	query := fmt.Sprintf(`
	SELECT COUNT(*) OVER(), id, article_id, user_id, body, created_at 
	FROM comments
	WHERE article_id = $1 and (to_tsvector('simple', body) @@ plainto_tsquery('simple', $2) OR $2 = '')
	ORDER BY %s %s limit $3 offset $4`, filters.sortColumn(), filters.sortDirection())

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := c.DB.QueryContext(ctx, query, article_id, body, filters.limit(), filters.offset())
	if err != nil {
		return nil, Metadata{}, err
	}

	defer rows.Close()

	totalRecords := 0
	comments := []*Comment{}

	for rows.Next() {
		var comment Comment

		err := rows.Scan(
			&totalRecords,
			&comment.Id,
			&comment.ArticleId,
			&comment.UserId,
			&comment.Body,
			&comment.CreatedAt,
		)
		if err != nil {
			return nil, Metadata{}, err
		}
		comments = append(comments, &comment)
	}
	// When the rows.Next() loop has finished, call rows.Err() to retrieve any error
	// that was encountered during the iteration.
	if err = rows.Err(); err != nil {
		return nil, Metadata{}, err
	}

	metadata := calculateMetadata(totalRecords, filters.Page, filters.Page_size)

	return comments, metadata, nil
}

func (c CommentModel) Insert(comment *Comment) error {

	query := `
		INSERT INTO comments(article_id, user_id, body) 
		VALUES ($1, $2, $3)
		RETURNING id, created_at
	`
	args := []interface{}{comment.ArticleId, comment.UserId, comment.Body}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return c.DB.QueryRowContext(ctx, query, args...).Scan(&comment.Id, &comment.CreatedAt)
}

func (c CommentModel) Get(comment_id int) (*Comment, error) {
	query := `
		SELECT id, article_id, user_id, body, created_at
		FROM comments
		WHERE id = $1
		`

	var comment Comment
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	row := c.DB.QueryRowContext(ctx, query, comment_id)

	err := row.Scan(&comment.Id, &comment.ArticleId, &comment.UserId, &comment.Body, &comment.CreatedAt)

	if err != nil {
		return nil, err
	}

	return &comment, nil
}

func (c *CommentModel) Update(comment *Comment) error {
	query := `
	UPDATE comments
	SET body = $1
	WHERE id= $2
	`
	args := []interface{}{comment.Body, comment.Id}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := c.DB.ExecContext(ctx, query, args...)

	return err
}

func (c CommentModel) Delete(id int) error {
	query := `
		DELETE FROM comments
		WHERE id = $1
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := c.DB.ExecContext(ctx, query, id)

	return err
}
