-- +goose Up

create table if not exists routes (
   id serial primary key,
   route varchar(1024) not null,
   method varchar(10) not null,
   description text not null,
   updated_at timestamp default now() not null,
   created_at timestamp not null default now(),
   deleted_at timestamp default null,
   namespace_id bigint references namespaces (id) not null,
   unique (route, method, namespace_id, deleted_at)
);

create unique index routes_method_route_deleted_at
    on routes (method, route, namespace_id, (deleted_at is null))
    where deleted_at is null;


-- +goose Down

drop table if exists routes;

