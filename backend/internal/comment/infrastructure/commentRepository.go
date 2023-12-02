package infrastructure

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/unmsmfisi-socialapplication/social_app/internal/comment/domain"
)


type CommentRepository struct {
    db *sql.DB 
}

func NewCommentRepository(database *sql.DB) *CommentRepository {
    return &CommentRepository{db: database}
}

func (r *CommentRepository) CreateComment(comment *domain.Comment) error {
    if comment == nil {
        return fmt.Errorf("comment cannot be nil")
    }

    query := `
        INSERT INTO SOC_APP_POSTS_COMMENTS (user_id, post_id, comment, insertion_date, update_date, parent_comment_id, is_active)
        VALUES ($1, $2, $3, $4, $5, $6, $7)
        RETURNING comment_id
    `
    err := r.db.QueryRow(
        query,
        comment.UserID,
        comment.PostID,
        comment.Comment,
        comment.InsertionDate,
        comment.UpdateDate,
        comment.ParentCommentID,
        comment.IsActive,
    ).Scan(&comment.CommentID)

    if err != nil {
        return fmt.Errorf("error creating comment: %w", err)
    }

    return nil
}

func (r *CommentRepository) GetAllComments() ([]*domain.Comment, error) {
    query := `
        SELECT comment_id, user_id, post_id, comment, insertion_date, update_date, parent_comment_id
        FROM SOC_APP_POSTS_COMMENTS
        WHERE is_active = true
    `
    rows, err := r.db.Query(query)

    if err != nil {
        return nil, fmt.Errorf("error querying all comments: %w", err)
    }

    defer rows.Close()

    comments := make([]*domain.Comment, 0)
    for rows.Next() {
        comment := &domain.Comment{}
        err := rows.Scan(
            &comment.CommentID,
            &comment.UserID,
            &comment.PostID,
            &comment.Comment,
            &comment.InsertionDate,
            &comment.UpdateDate,
            &comment.ParentCommentID,
        )

        if err != nil {
            return nil, fmt.Errorf("error scanning comment: %w", err)
        }

        comments = append(comments, comment)
    }

    if err = rows.Err(); err != nil {
        return nil, fmt.Errorf("error after iterating over rows: %w", err)
    }

    return comments, nil
}

func (r *CommentRepository) GetCommentByID(commentID int64) (*domain.Comment, error) {
    if commentID <= 0 {
        return nil, fmt.Errorf("invalid comment ID")
    }

    query := `
        SELECT comment_id, user_id, post_id, comment, insertion_date, update_date, parent_comment_id
        FROM SOC_APP_POSTS_COMMENTS
        WHERE comment_id = $1 and is_active = true
    `
    row := r.db.QueryRow(query, commentID)
    comment := &domain.Comment{}
    err := row.Scan(
        &comment.CommentID,
        &comment.UserID,
        &comment.PostID,
        &comment.Comment,
        &comment.InsertionDate,
        &comment.UpdateDate,
        &comment.ParentCommentID,
    )

    if err != nil {
        return nil, fmt.Errorf("error retrieving comment by ID: %w", err)
    }

    return comment, nil
}

func (r *CommentRepository) GetCommentsByPostID(postID int64) ([]*domain.Comment, error) {
    if postID <= 0 {
        return nil, fmt.Errorf("invalid post ID")
    }

    query := `
        SELECT comment_id, user_id, post_id, comment, insertion_date, update_date, parent_comment_id
        FROM SOC_APP_POSTS_COMMENTS
        WHERE post_id = $1 and is_active = true
    `
    rows, err := r.db.Query(query, postID)

    if err != nil {
        return nil, fmt.Errorf("error querying comments: %w", err)
    }

    defer rows.Close()

    comments := make([]*domain.Comment, 0)
    for rows.Next() {
        comment := &domain.Comment{}
        err := rows.Scan(
            &comment.CommentID,
            &comment.UserID,
            &comment.PostID,
            &comment.Comment,
            &comment.InsertionDate,
            &comment.UpdateDate,
            &comment.ParentCommentID,
        )

        if err != nil {
            return nil, err
        }

        comments = append(comments, comment)
    }

    if err = rows.Err(); err != nil {
        return nil, fmt.Errorf("error after iterating over rows: %w", err)
    }

    return comments, nil
}

func (r *CommentRepository) UpdateComment(commentID int64, newComment *domain.Comment) error {
    if newComment == nil || commentID <= 0 {
        return fmt.Errorf("invalid input")
    }

    currentComment, err := r.GetCommentByID(commentID)
    if err != nil {
        return fmt.Errorf("error fetching current comment: %w", err)
    }

    var queryBuilder strings.Builder
    queryBuilder.WriteString("UPDATE SOC_APP_POSTS_COMMENTS SET ")
    updateFields := make([]interface{}, 0)
    fieldCount := 1

    if newComment.UserID != currentComment.UserID {
        queryBuilder.WriteString(fmt.Sprintf("user_id = $%d, ", fieldCount))
        updateFields = append(updateFields, newComment.UserID)
        fieldCount++
    }

    query := strings.TrimSuffix(queryBuilder.String(), ", ")
    query += fmt.Sprintf(" WHERE comment_id = $%d and is_active = true", fieldCount)
    updateFields = append(updateFields, commentID)

    result, err := r.db.Exec(query, updateFields...)
    if err != nil {
        return fmt.Errorf("error updating comment: %w", err)
    }

    rowsAffected, err := result.RowsAffected()
    if err != nil {
        return fmt.Errorf("error checking affected rows: %w", err)
    }
    if rowsAffected == 0 {
        return fmt.Errorf("no rows were updated")
    }

    return nil
}

func (r *CommentRepository) DeleteComment(commentID int64) error {
    if commentID <= 0 {
        return fmt.Errorf("invalid comment ID")
    }

    query := `
        UPDATE SOC_APP_POSTS_COMMENTS
        SET is_active = false
        WHERE comment_id = $1
    `

    result, err := r.db.Exec(query, commentID)
    if err != nil {
        return fmt.Errorf("error deleting comment: %w", err)
    }

    rowsAffected, err := result.RowsAffected()
    if err != nil {
        return fmt.Errorf("error checking affected rows: %w", err)
    }
    if rowsAffected == 0 {
        return fmt.Errorf("no rows were updated")
    }

    return nil
}