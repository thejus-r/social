package store

import (
	"context"
	"database/sql"
	"errors"

	pg "github.com/lib/pq"
)

type Post struct {
	ID        int64     `json:"id"`
	Content   string    `json:"content"`
	Title     string    `json:"title"`
	UserId    int64     `json:"user_ID"`
	Tags      []string  `json:"tags"`
	CreatedAt string    `json:"created_at"`
	UpdatedAt string    `json:"updated_at"`
	Comments  []Comment `json:"comments"`
	User      User      `json:"user"`
}

type PostWithMetadata struct {
	Post
	CommentCount int `json:"comment_count"`
}

type PostStore struct {
	db *sql.DB
}

type UpdatePostPayload struct {
}

func (s *PostStore) GetUserFeed(ctx context.Context, userID int64, fq PaginatedFeedQuery) ([]PostWithMetadata, error) {

	query := `
		SELECT p.id, p.user_id, p.title, p.content, p.created_at, p.tags, u.username,COUNT(c.id) AS comments_count
		FROM posts p
		LEFT JOIN comments c ON c.post_id = p.id
		LEFT JOIN users u ON p.user_id = u.id
		JOIN followers f ON f.follower_id = p.user_id OR p.user_id = $1 
		WHERE
			f.user_id = $1 AND
			(p.title ILIKE '%' || $4 || '%' OR p.content ILIKE '%' || $4 || '%') AND
			(p.tags @> $5 OR $5 = '{}')
		GROUP BY p.id, u.username
		ORDER BY p.created_at ` + fq.Sort + ` 
		LIMIT $2 OFFSET $3;
	`
	rows, err := s.db.QueryContext(ctx, query, userID, fq.Limit, fq.Offset, fq.Search, pg.Array(fq.Tags))
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var feed []PostWithMetadata
	for rows.Next() {
		var p PostWithMetadata
		err := rows.Scan(
			&p.ID,
			&p.UserId,
			&p.Title,
			&p.Content,
			&p.CreatedAt,
			pg.Array(&p.Tags),
			&p.User.Username,
			&p.CommentCount,
		)
		if err != nil {
			return nil, err
		}

		feed = append(feed, p)
	}

	return feed, nil
}

func (s *PostStore) Update(ctx context.Context, post *Post) error {

	query := `
	UPDATE posts
	SET title = $1, content = $2
	WHERE id = $3
	`

	_, err := s.db.ExecContext(
		ctx,
		query,
		post.Title,
		post.Content,
		post.ID,
	)

	if err != nil {
		return err
	}

	return nil
}

func (s *PostStore) Create(ctx context.Context, post *Post) error {

	query := `
	INSERT INTO posts (content, title, user_id, tags)
	VALUES ($1, $2, $3, $4) RETURNING id, created_at, updated_at
	`
	err := s.db.QueryRowContext(
		ctx,
		query,
		post.Content,
		post.Title,
		post.UserId,
		pg.Array(post.Tags),
	).Scan(
		&post.ID,
		&post.CreatedAt,
		&post.UpdatedAt,
	)
	if err != nil {
		return err
	}
	return nil
}

func (s *PostStore) GetByID(ctx context.Context, id int64) (*Post, error) {

	query := `
	SELECT id, user_id, title, content, created_at, updated_at, tags
	FROM posts
	WHERE id = $1
	`

	var post Post

	err := s.db.QueryRowContext(
		ctx,
		query,
		id,
	).Scan(
		&post.ID,
		&post.UserId,
		&post.Title,
		&post.Content,
		&post.CreatedAt,
		&post.UpdatedAt,
		pg.Array(&post.Tags),
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrNotFound
		default:
			return nil, err
		}
	}
	return &post, nil
}

func (s *PostStore) Delete(ctx context.Context, postID int64) error {

	query := `DELETE FROM posts WHERE id = $1`

	res, err := s.db.ExecContext(ctx, query, postID)
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return ErrNotFound
	}

	return nil
}
