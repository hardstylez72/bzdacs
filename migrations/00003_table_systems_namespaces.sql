-- +goose Up

create table if not exists systems_namespaces (
  system_id int references systems (id) not null,
  namespace_id bigint references namespaces (id) not null,
  unique (system_id, namespace_id)
);

-- +goose Down

drop table if exists systems_namespaces;