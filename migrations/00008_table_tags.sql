-- +goose Up

create table if not exists ad.tags (
   id serial primary key,
   name text not null,
   updated_at timestamp default null,
   created_at timestamp not null default now(),
   deleted_at timestamp default null,
   unique (deleted_at, name)
);

-- +goose Down

drop table if exists ad.tags;