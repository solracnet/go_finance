-- name: CreateCategory :one
insert into categories (user_id, title, type, description) values ($1, $2, $3, $4) returning *;

-- name: GetCategoryById :one
select * from categories where id = $1 limit 1;

-- name: GetCategories :many
select * from categories where user_id = $1 and type = $2 and title ilike concat('%', @title::text, '%') and description ilike concat('%', sqlc.arg('description')::text, '%');

-- name: GetCategoriesByUserIdAndType :many
select * from categories where user_id = $1 and type = $2;

-- name: GetCategoriesByUserIdAndTypeAndTitle :many
select * from categories where user_id = $1 and type = $2 and title like $3;

-- name: GetCategoriesByUserIdAndTypeAndDescription :many
select * from categories where user_id = $1 and type = $2 and description like $3;

-- name: UpdateCategory :one
update categories set title = $2, description = $3 where id = $1 returning *;

-- name: DeleteCategory :exec
delete from categories where id = $1;