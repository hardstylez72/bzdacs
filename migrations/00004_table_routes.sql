-- +goose Up

create table if not exists ad.routes (
   id serial primary key,
   route varchar(1024) not null,
   method varchar(10) not null,
   description text not null,
   updated_at timestamp default null,
   created_at timestamp not null default now(),
   deleted_at timestamp default null,
   unique (route, method, deleted_at)
);

create unique index routes_method_route_deleted_at
    on ad.routes (method, route, (deleted_at is null))
    where deleted_at is null;


-- +goose Down

drop table if exists ad.routes;

