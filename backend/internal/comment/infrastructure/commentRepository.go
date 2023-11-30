package infrastructure

import (
	"database/sql"

	"github.com/unmsmfisi-socialapplication/social_app/internal/comment/domain"
)


type CommentRepository struct {
    db *sql.DB 
}

func NewCommentRepository(database *sql.DB) *CommentRepository {
    return &CommentRepository{db: database}
}

// Code to create a comment in the database
func (r *CommentRepository) CreateComment(comment *domain.Comment) error {
    query := `
        INSERT INTO SOC_APP_POSTS_COMMENTS (user_id, post_id, comment, insertion_date, update_date, parent_comment_id)
        VALUES ($1, $2, $3, $4, $5, $6)
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
    ).Scan(&comment.CommentID)

    if err != nil {
        return err
    }

    return nil
}

func (r *CommentRepository) GetAllComments() ([]*domain.Comment, error) {
    query := `
        SELECT comment_id, user_id, post_id, comment, insertion_date, update_date, parent_comment_id
        FROM SOC_APP_POSTS_COMMENTS
    `
    rows, err := r.db.Query(query)

    if err != nil {
        return nil, err
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
        return nil, err
    }

    return comments, nil
}

// Code to get a comment by its ID
func (r *CommentRepository) GetCommentByID(commentID int64) (*domain.Comment, error) {
    query := `
        SELECT comment_id, user_id, post_id, comment, insertion_date, update_date, parent_comment_id
        FROM SOC_APP_POSTS_COMMENTS
        WHERE comment_id = $1
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
        return nil, err
    }

    return comment, nil
}

func (r *CommentRepository) GetCommentsByPostID(postID int64) ([]*domain.Comment, error) {
    query := `
        SELECT comment_id, user_id, post_id, comment, insertion_date, update_date, parent_comment_id
        FROM SOC_APP_POSTS_COMMENTS
        WHERE post_id = $1
    `
    rows, err := r.db.Query(query, postID)

    if err != nil {
        return nil, err
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
        return nil, err
    }

    return comments, nil
}

//Code to update a comment in the database
func (r *CommentRepository) UpdateComment(commentID int64, comment *domain.Comment) error {
    query := `
        UPDATE SOC_APP_POSTS_COMMENTS
        SET user_id = $2, post_id = $3, comment = $4, update_date = $5, parent_comment_id = $6
        WHERE comment_id = $1
    `
    _, err := r.db.Exec(
        query,
        commentID,
        comment.UserID,
        comment.PostID,
        comment.Comment,
        comment.UpdateDate,
        comment.ParentCommentID,
    )

    if err != nil {
        return err
    }

    return nil
}

//Code to delete a comment in the database
func (r *CommentRepository) DeleteComment(commentID int64) error {
    query := `
        DELETE FROM SOC_APP_POSTS_COMMENTS
        WHERE comment_id = $1
    `
    _, err := r.db.Exec(query, commentID)

    if err != nil {
        return err
    }

    return nil
}