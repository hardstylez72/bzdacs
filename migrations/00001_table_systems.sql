-- +goose Up

create table if not exists systems (
    id serial primary key,
    name text not null,
    updated_at timestamp not null default now(),
    created_at timestamp not null default now(),
    deleted_at timestamp default null,
    unique (name, deleted_at)
);
create unique index systems_name_deleted_at
    on systems (name, (deleted_at is null))
    where deleted_at is null;

-- +goose Down

drop table if exists systems;