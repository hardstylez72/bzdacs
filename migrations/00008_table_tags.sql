-- +goose Up

create table if not exists tags (
   id serial primary key,
   name text not null,
   updated_at timestamp default now() not null,
   created_at timestamp not null default now(),
   deleted_at timestamp default null,
   unique (name, deleted_at)
);
create unique index tags_name_deleted_at
    on tags (name, (deleted_at is null))
    where deleted_at is null;

-- +goose Down

drop table if exists tags;