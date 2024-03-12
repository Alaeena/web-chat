-- name: CreateUser :one
insert into users (username, email, password)
values ($1, $2, $3)
returning *;

-- name: GetUserByEmail :one
select * from users where email = $1;