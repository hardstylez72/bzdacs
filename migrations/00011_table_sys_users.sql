-- +goose Up

create table if not exists sys_users (
    id bigserial primary key,
    login varchar(4095) not null,
    password varchar(4095) not null,
    created_at timestamp not null default now(),
    updated_at timestamp not null default now(),
    deleted_at timestamp default null,
    unique (id, login)
);

-- +goose Down

drop table if exists sys_users;