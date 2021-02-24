-- +goose Up

create table if not exists users (
    id bigserial primary key,
    external_id text not null,
    created_at timestamp not null default now(),
    updated_at timestamp not null default now(),
    deleted_at timestamp default null,
    namespace_id bigint references namespaces (id) not null,
    unique (external_id, namespace_id, deleted_at)
);

create unique index users_external_id_deleted_at
    on users (external_id, namespace_id, (deleted_at is null))
    where deleted_at is null;

-- +goose Down

drop table if exists users;