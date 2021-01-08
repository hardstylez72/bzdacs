-- +goose Up

create table if not exists ad.users_groups (
   group_id int references ad.groups (id),
   user_id bigint references ad.users (id),
   unique (group_id, user_id)
);

-- +goose Down

drop table if exists ad.users_groups;