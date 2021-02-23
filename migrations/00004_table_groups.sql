-- +goose Up

create table if not exists groups (
     id serial primary key,
     code text not null,
     description text not null,
     created_at timestamp default now() not null,
     updated_at timestamp default now() not null,
     deleted_at timestamp default null,
     namespace_id bigint references namespaces (id) not null,
     unique (code, namespace_id, deleted_at)
);

create unique index groups_code_deleted_at
    on groups (code, namespace_id, (deleted_at is null))
    where deleted_at is null;

-- +goose Down

drop table if exists groups;