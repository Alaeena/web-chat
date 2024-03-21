-- name: CreateUser :one
insert into users (username, email, password)
values ($1, $2, $3)
returning *;

-- name: GetUserByEmail :one
select * from users where email = $1;

-- name: GetUserByTokenAndId :one
select * from users where access_token = $1 and id = $2;

-- name: SetAccessToken :exec
update users set access_token = $1 where id = $2;