-- +goose Up

create table if not exists ad.routes_tags (
   route_id int references ad.routes (id),
   tag_id bigint references ad.tags (id),
   unique (route_id, tag_id)
);

-- +goose Down

drop table if exists ad.routes_tags;