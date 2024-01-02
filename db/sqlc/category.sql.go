// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: category.sql

package db

import (
	"context"
)

const createCategory = `-- name: CreateCategory :one
insert into categories (user_id, title, type, description) values ($1, $2, $3, $4) returning id, user_id, title, type, description, created_at, updated_at
`

type CreateCategoryParams struct {
	UserID      int32  `json:"user_id"`
	Title       string `json:"title"`
	Type        string `json:"type"`
	Description string `json:"description"`
}

func (q *Queries) CreateCategory(ctx context.Context, arg CreateCategoryParams) (Category, error) {
	row := q.queryRow(ctx, q.createCategoryStmt, createCategory,
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
		&i.UpdatedAt,
	)
	return i, err
}

const deleteCategory = `-- name: DeleteCategory :exec
delete from categories where id = $1
`

func (q *Queries) DeleteCategory(ctx context.Context, id int32) error {
	_, err := q.exec(ctx, q.deleteCategoryStmt, deleteCategory, id)
	return err
}

const getCategories = `-- name: GetCategories :many
select id, user_id, title, type, description, created_at, updated_at from categories where user_id = $1 and type = $2 and title ilike concat('%', $3::text, '%') and description ilike concat('%', $4::text, '%')
`

type GetCategoriesParams struct {
	UserID      int32  `json:"user_id"`
	Type        string `json:"type"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (q *Queries) GetCategories(ctx context.Context, arg GetCategoriesParams) ([]Category, error) {
	rows, err := q.query(ctx, q.getCategoriesStmt, getCategories,
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
			&i.UpdatedAt,
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

const getCategoriesByUserIdAndType = `-- name: GetCategoriesByUserIdAndType :many
select id, user_id, title, type, description, created_at, updated_at from categories where user_id = $1 and type = $2
`

type GetCategoriesByUserIdAndTypeParams struct {
	UserID int32  `json:"user_id"`
	Type   string `json:"type"`
}

func (q *Queries) GetCategoriesByUserIdAndType(ctx context.Context, arg GetCategoriesByUserIdAndTypeParams) ([]Category, error) {
	rows, err := q.query(ctx, q.getCategoriesByUserIdAndTypeStmt, getCategoriesByUserIdAndType, arg.UserID, arg.Type)
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
			&i.UpdatedAt,
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

const getCategoriesByUserIdAndTypeAndDescription = `-- name: GetCategoriesByUserIdAndTypeAndDescription :many
select id, user_id, title, type, description, created_at, updated_at from categories where user_id = $1 and type = $2 and description like $3
`

type GetCategoriesByUserIdAndTypeAndDescriptionParams struct {
	UserID      int32  `json:"user_id"`
	Type        string `json:"type"`
	Description string `json:"description"`
}

func (q *Queries) GetCategoriesByUserIdAndTypeAndDescription(ctx context.Context, arg GetCategoriesByUserIdAndTypeAndDescriptionParams) ([]Category, error) {
	rows, err := q.query(ctx, q.getCategoriesByUserIdAndTypeAndDescriptionStmt, getCategoriesByUserIdAndTypeAndDescription, arg.UserID, arg.Type, arg.Description)
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
			&i.UpdatedAt,
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

const getCategoriesByUserIdAndTypeAndTitle = `-- name: GetCategoriesByUserIdAndTypeAndTitle :many
select id, user_id, title, type, description, created_at, updated_at from categories where user_id = $1 and type = $2 and title like $3
`

type GetCategoriesByUserIdAndTypeAndTitleParams struct {
	UserID int32  `json:"user_id"`
	Type   string `json:"type"`
	Title  string `json:"title"`
}

func (q *Queries) GetCategoriesByUserIdAndTypeAndTitle(ctx context.Context, arg GetCategoriesByUserIdAndTypeAndTitleParams) ([]Category, error) {
	rows, err := q.query(ctx, q.getCategoriesByUserIdAndTypeAndTitleStmt, getCategoriesByUserIdAndTypeAndTitle, arg.UserID, arg.Type, arg.Title)
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
			&i.UpdatedAt,
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

const getCategoryById = `-- name: GetCategoryById :one
select id, user_id, title, type, description, created_at, updated_at from categories where id = $1 limit 1
`

func (q *Queries) GetCategoryById(ctx context.Context, id int32) (Category, error) {
	row := q.queryRow(ctx, q.getCategoryByIdStmt, getCategoryById, id)
	var i Category
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Title,
		&i.Type,
		&i.Description,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateCategory = `-- name: UpdateCategory :one
update categories set title = $2, description = $3 where id = $1 returning id, user_id, title, type, description, created_at, updated_at
`

type UpdateCategoryParams struct {
	ID          int32  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (q *Queries) UpdateCategory(ctx context.Context, arg UpdateCategoryParams) (Category, error) {
	row := q.queryRow(ctx, q.updateCategoryStmt, updateCategory, arg.ID, arg.Title, arg.Description)
	var i Category
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Title,
		&i.Type,
		&i.Description,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
