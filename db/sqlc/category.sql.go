// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: category.sql

package db

import (
	"context"
)

const createCategory = `-- name: CreateCategory :one
INSERT INTO CATEGORIES (
  USER_ID,
  TITLE,
  TYPE,
  DESCRIPTION
) VALUES (
  $1,
  $2,
  $3,
  $4
) RETURNING id, user_id, title, type, description, created_at
`

type CreateCategoryParams struct {
	UserID      int32  `json:"user_id"`
	Title       string `json:"title"`
	Type        string `json:"type"`
	Description string `json:"description"`
}

func (q *Queries) CreateCategory(ctx context.Context, arg CreateCategoryParams) (Category, error) {
	row := q.db.QueryRowContext(ctx, createCategory,
		arg.UserID,
		arg.Title,
		arg.Type,
		arg.Description,
	)
	var i Category
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Title,
		&i.Type,
		&i.Description,
		&i.CreatedAt,
	)
	return i, err
}

const deleteCategories = `-- name: DeleteCategories :exec
DELETE FROM CATEGORIES
WHERE
  ID = $1
`

func (q *Queries) DeleteCategories(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteCategories, id)
	return err
}

const getCategories = `-- name: GetCategories :many
SELECT
  id, user_id, title, type, description, created_at
FROM
  CATEGORIES
WHERE
  USER_ID = $1
  AND TYPE = $2
  AND LOWER(TITLE) LIKE CONCAT('%',
  LOWER($3::TEXT),
  '%')
  AND LOWER(DESCRIPTION) LIKE CONCAT('%',
  LOWER($4::TEXT),
  '%')
`

type GetCategoriesParams struct {
	UserID      int32  `json:"user_id"`
	Type        string `json:"type"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (q *Queries) GetCategories(ctx context.Context, arg GetCategoriesParams) ([]Category, error) {
	rows, err := q.db.QueryContext(ctx, getCategories,
		arg.UserID,
		arg.Type,
		arg.Title,
		arg.Description,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Category{}
	for rows.Next() {
		var i Category
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Title,
			&i.Type,
			&i.Description,
			&i.CreatedAt,
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

const getCategory = `-- name: GetCategory :one
SELECT
  id, user_id, title, type, description, created_at
FROM
  CATEGORIES
WHERE
  ID = $1 LIMIT 1
`

func (q *Queries) GetCategory(ctx context.Context, id int32) (Category, error) {
	row := q.db.QueryRowContext(ctx, getCategory, id)
	var i Category
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Title,
		&i.Type,
		&i.Description,
		&i.CreatedAt,
	)
	return i, err
}

const updateCategories = `-- name: UpdateCategories :one
UPDATE CATEGORIES
SET
  TITLE = $2,
  DESCRIPTION = $3
WHERE
  ID = $1 RETURNING id, user_id, title, type, description, created_at
`

type UpdateCategoriesParams struct {
	ID          int32  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (q *Queries) UpdateCategories(ctx context.Context, arg UpdateCategoriesParams) (Category, error) {
	row := q.db.QueryRowContext(ctx, updateCategories, arg.ID, arg.Title, arg.Description)
	var i Category
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Title,
		&i.Type,
		&i.Description,
		&i.CreatedAt,
	)
	return i, err
}

const getCategoriesByUserIdAndType = `-- name: getCategoriesByUserIdAndType :many
SELECT
  id, user_id, title, type, description, created_at
FROM
  CATEGORIES
WHERE
  USER_ID = $1
  AND TYPE = $2
`

type getCategoriesByUserIdAndTypeParams struct {
	UserID int32  `json:"user_id"`
	Type   string `json:"type"`
}

func (q *Queries) getCategoriesByUserIdAndType(ctx context.Context, arg getCategoriesByUserIdAndTypeParams) ([]Category, error) {
	rows, err := q.db.QueryContext(ctx, getCategoriesByUserIdAndType, arg.UserID, arg.Type)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Category{}
	for rows.Next() {
		var i Category
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Title,
			&i.Type,
			&i.Description,
			&i.CreatedAt,
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

const getCategoriesByUserIdAndTypeAndDescription = `-- name: getCategoriesByUserIdAndTypeAndDescription :many

SELECT
  id, user_id, title, type, description, created_at
FROM
  CATEGORIES
WHERE
  USER_ID = $1
  AND TYPE = $2
  AND LOWER(DESCRIPTION) LIKE CONCAT('%',
  LOWER($3::TEXT),
  '%')
`

type getCategoriesByUserIdAndTypeAndDescriptionParams struct {
	UserID      int32  `json:"user_id"`
	Type        string `json:"type"`
	Description string `json:"description"`
}

func (q *Queries) getCategoriesByUserIdAndTypeAndDescription(ctx context.Context, arg getCategoriesByUserIdAndTypeAndDescriptionParams) ([]Category, error) {
	rows, err := q.db.QueryContext(ctx, getCategoriesByUserIdAndTypeAndDescription, arg.UserID, arg.Type, arg.Description)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Category{}
	for rows.Next() {
		var i Category
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Title,
			&i.Type,
			&i.Description,
			&i.CreatedAt,
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

const getCategoriesByUserIdAndTypeAndTitle = `-- name: getCategoriesByUserIdAndTypeAndTitle :many
SELECT
  id, user_id, title, type, description, created_at
FROM
  CATEGORIES
WHERE
  USER_ID = $1
  AND TYPE = $2
  AND LOWER(TITLE) LIKE CONCAT('%',
  LOWER($3::TEXT),
  '%')
`

type getCategoriesByUserIdAndTypeAndTitleParams struct {
	UserID int32  `json:"user_id"`
	Type   string `json:"type"`
	Title  string `json:"title"`
}

func (q *Queries) getCategoriesByUserIdAndTypeAndTitle(ctx context.Context, arg getCategoriesByUserIdAndTypeAndTitleParams) ([]Category, error) {
	rows, err := q.db.QueryContext(ctx, getCategoriesByUserIdAndTypeAndTitle, arg.UserID, arg.Type, arg.Title)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Category{}
	for rows.Next() {
		var i Category
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Title,
			&i.Type,
			&i.Description,
			&i.CreatedAt,
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
