-- +goose Up

create table if not exists ad.users (
    id bigserial primary key,
    external_id text not null,

    created_at timestamp default now() not null,
    updated_at timestamp default null,
    deleted_at timestamp default null,
    unique (external_id, deleted_at)
);

create unique index users_external_id_deleted_at
    on ad.users (external_id, (deleted_at is null))
    where deleted_at is null;

-- +goose Down

drop table if exists ad.users;