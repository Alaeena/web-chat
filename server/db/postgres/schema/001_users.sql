-- +goose Up
create table users(
    id serial primary key,
    username text not null ,
    email text not null unique,
    password text not null,
    access_token text,
    refresh_token text
);

-- +goose Down

drop table users;