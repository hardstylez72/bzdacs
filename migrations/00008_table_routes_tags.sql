-- +goose Up

create table if not exists routes_tags (
   route_id int references routes (id)  not null,
   tag_id bigint references tags (id)  not null,
   unique (route_id, tag_id)
);

-- +goose Down

drop table if exists routes_tags;