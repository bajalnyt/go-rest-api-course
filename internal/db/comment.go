package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/bajalnyt/go-rest-api-course/internal/comment"
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

func (d *Database) GetComment(
	ctx context.Context,
	uuid string,
) (comment.Comment, error) {
	var cmtRow CommentRow
	row := d.Client.QueryRowContext(ctx, `SELECT id ,slug, author, body from comments where id=$1`, uuid)
	err := row.Scan(&cmtRow.ID, &cmtRow.Slug, &cmtRow.Author, &cmtRow.Body)
	if err != nil {
		return comment.Comment{}, fmt.Errorf("error fetching comment by uuid %w", err)
	}

	return comment.Comment{}, nil
}
