-- +goose Up

create table if not exists groups_routes (
   route_id int  references routes (id) not null,
   group_id int references groups (id) not null,
   UNIQUE  (route_id, group_id)
);

-- +goose Down

drop table if exists groups_routes;