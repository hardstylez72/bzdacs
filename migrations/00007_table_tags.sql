-- +goose Up

create table if not exists tags (
   id serial primary key,
   name text not null,
   updated_at timestamp default now() not null,
   created_at timestamp default now() not null,
   deleted_at timestamp default null,
   namespace_id bigint references namespaces (id) not null,
   unique (name, namespace_id, deleted_at)
);
create unique index tags_name_deleted_at
    on tags (name, namespace_id, (deleted_at is null))
    where deleted_at is null;

-- +goose Down

drop table if exists tags;