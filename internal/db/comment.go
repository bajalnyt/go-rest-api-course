package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/bajalnyt/go-rest-api-course/internal/comment"
	"github.com/google/uuid"
)

type CommentRow struct {
	ID     string
	Slug   sql.NullString
	Body   sql.NullString
	Author sql.NullString
}

func convertCommentRowToComment(c CommentRow) comment.Comment {
	return comment.Comment{
		ID:     c.ID,
		Slug:   c.Slug.String,
		Author: c.Author.String,
		Body:   c.Body.String,
	}

}

func (d *Database) GetComments(
	ctx context.Context,
) ([]comment.Comment, error) {
	var cmtRow CommentRow
	var comments []comment.Comment
	rows, err := d.Client.Query(`SELECT id ,slug, author, body from comments`)
	if err != nil {
		return []comment.Comment{}, fmt.Errorf("error fetching comments %w", err)
	}
	defer rows.Next()

	for rows.Next() {
		if err := rows.Scan(&cmtRow.ID, &cmtRow.Slug, &cmtRow.Author, &cmtRow.Body); err != nil {
			log.Fatal(err)
		}
		comments = append(comments, convertCommentRowToComment(cmtRow))
	}

	return comments, nil
}

func (d *Database) GetComment(
	ctx context.Context,
	uuid string,
) (comment.Comment, error) {
	var cmtRow CommentRow
	row := d.Client.QueryRowContext(ctx,
		`SELECT id ,slug, author, body from comments where id=$1`,
		uuid)
	err := row.Scan(&cmtRow.ID, &cmtRow.Slug, &cmtRow.Author, &cmtRow.Body)
	if err != nil {
		return comment.Comment{}, fmt.Errorf("error fetching comment by uuid %w", err)
	}

	return convertCommentRowToComment(cmtRow), nil
}

func (d *Database) PostComment(
	ctx context.Context,
	cmt comment.Comment,
) (comment.Comment, error) {
	cmt.ID = uuid.New().String()
	postRow := CommentRow{
		ID:     cmt.ID,
		Slug:   sql.NullString{String: cmt.Slug, Valid: true},
		Author: sql.NullString{String: cmt.Author, Valid: true},
		Body:   sql.NullString{String: cmt.Body, Valid: true},
	}

	rows, err := d.Client.NamedQueryContext(
		ctx,
		`INSERT INTO comments
		(id, slug, author, body)
		values
		(:id, :slug, :author, :body)`,
		postRow,
	)
	if err != nil {
		return comment.Comment{}, err
	}

	if err := rows.Close(); err != nil {
		return comment.Comment{}, err
	}

	return convertCommentRowToComment(postRow), nil
}

// UpdateComment - updates a comment in the database
func (d *Database) UpdateComment(ctx context.Context, id string, cmt comment.Comment) (comment.Comment, error) {
	cmtRow := CommentRow{
		ID:     id,
		Slug:   sql.NullString{String: cmt.Slug, Valid: true},
		Body:   sql.NullString{String: cmt.Body, Valid: true},
		Author: sql.NullString{String: cmt.Author, Valid: true},
	}

	rows, err := d.Client.NamedQueryContext(
		ctx,
		`UPDATE comments SET
		slug = :slug,
		author = :author,
		body = :body 
		WHERE id = :id`,
		cmtRow,
	)
	if err != nil {
		return comment.Comment{}, fmt.Errorf("failed to insert comment: %w", err)
	}
	if err := rows.Close(); err != nil {
		return comment.Comment{}, fmt.Errorf("failed to close rows: %w", err)
	}

	return convertCommentRowToComment(cmtRow), nil
}

// DeleteComment - deletes a comment from the database
func (d *Database) DeleteComment(ctx context.Context, id string) error {
	_, err := d.Client.ExecContext(
		ctx,
		`DELETE FROM comments where id = $1`,
		id,
	)
	if err != nil {
		return fmt.Errorf("failed to delete comment from the database: %w", err)
	}
	return nil
}
