-- +goose Up

create table if not exists users_routes (
   route_id int references routes (id) not null,
   user_id bigint references users (id) not null,
   is_excluded bool default false not null,
   unique (route_id, user_id)
);

-- +goose Down

drop table if exists users_routes;