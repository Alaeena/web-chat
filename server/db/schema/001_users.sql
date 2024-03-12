-- +goose Up
create table users(
    id serial primary key,
    username text not null ,
    email text not null unique,
    password text not null
);

-- +goose Down

drop table users;