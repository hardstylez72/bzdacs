-- +goose Up

create table if not exists users_groups (
   group_id int references groups (id) not null,
   user_id bigint references users (id) not null,
   unique (group_id, user_id)
);

-- +goose Down

drop table if exists users_groups;