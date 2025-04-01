// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: feed_follows.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createFeedFollow = `-- name: CreateFeedFollow :many
WITH inserted_feed_follows AS (
    INSERT INTO feed_follows (id, created_at, updated_at, feed_id, user_id)
    VALUES (
        $1,
        $2,
        $3,
        $4,
        $5
    )
    RETURNING id, created_at, updated_at, feed_id, user_id
) SELECT inserted_feed_follows.id, inserted_feed_follows.created_at, inserted_feed_follows.updated_at, inserted_feed_follows.feed_id, inserted_feed_follows.user_id,
        feeds.name AS feeds_name,
        users.name AS users_name
FROM inserted_feed_follows
INNER JOIN feeds
ON feeds.id = inserted_feed_follows.feed_id
INNER JOIN users
ON users.id = inserted_feed_follows.user_id
`

type CreateFeedFollowParams struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	FeedID    uuid.UUID
	UserID    uuid.UUID
}

type CreateFeedFollowRow struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	FeedID    uuid.UUID
	UserID    uuid.UUID
	FeedsName string
	UsersName string
}

func (q *Queries) CreateFeedFollow(ctx context.Context, arg CreateFeedFollowParams) ([]CreateFeedFollowRow, error) {
	rows, err := q.db.QueryContext(ctx, createFeedFollow,
		arg.ID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.FeedID,
		arg.UserID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []CreateFeedFollowRow
	for rows.Next() {
		var i CreateFeedFollowRow
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.FeedID,
			&i.UserID,
			&i.FeedsName,
			&i.UsersName,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const deleteFeedFollowsForUser = `-- name: DeleteFeedFollowsForUser :exec
DELETE FROM feed_follows
WHERE user_id = $1 AND feed_id = $2
`

type DeleteFeedFollowsForUserParams struct {
	UserID uuid.UUID
	FeedID uuid.UUID
}

func (q *Queries) DeleteFeedFollowsForUser(ctx context.Context, arg DeleteFeedFollowsForUserParams) error {
	_, err := q.db.ExecContext(ctx, deleteFeedFollowsForUser, arg.UserID, arg.FeedID)
	return err
}

const getFeedFollowsForUser = `-- name: GetFeedFollowsForUser :many
SELECT feed_follows.id, feed_follows.created_at, feed_follows.updated_at, feed_follows.feed_id, feed_follows.user_id, feeds.name AS feeds_name, users.name AS users_name
FROM feed_follows
INNER JOIN feeds
ON feeds.id = feed_follows.feed_id
INNER JOIN users
ON users.id = feed_follows.user_id
WHERE feed_follows.user_id = $1
`

type GetFeedFollowsForUserRow struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	FeedID    uuid.UUID
	UserID    uuid.UUID
	FeedsName string
	UsersName string
}

func (q *Queries) GetFeedFollowsForUser(ctx context.Context, userID uuid.UUID) ([]GetFeedFollowsForUserRow, error) {
	rows, err := q.db.QueryContext(ctx, getFeedFollowsForUser, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetFeedFollowsForUserRow
	for rows.Next() {
		var i GetFeedFollowsForUserRow
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.FeedID,
			&i.UserID,
			&i.FeedsName,
			&i.UsersName,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
