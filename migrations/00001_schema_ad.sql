-- +goose Up

create schema if not exists ad;

-- +goose Down

DROP schema if exists ad;