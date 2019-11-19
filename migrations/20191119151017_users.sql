-- +goose Up
-- SQL in this section is executed when the migration is applied.
create table if not exists users(
    id SERIAL PRIMARY KEY,
    user_id UUID UNIQUE not null,
    username varchar(100) not null,
    password varchar(100) not null
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
drop table if exists users;
