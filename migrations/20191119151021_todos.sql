-- +goose Up
-- SQL in this section is executed when the migration is applied.
create table if not exists todos (
    id serial primary key,
    title varchar(100) not null default '',
    text text not null  default '',
    user_id bigint not null references users (id) on delete CASCADE
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
drop table if exists todos
