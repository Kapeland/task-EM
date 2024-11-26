-- +goose Up
-- +goose StatementBegin
drop schema if exists library_schema;

create schema if not exists library_schema;

create table if not exists library_schema.library (
    song_group text,
    song text,
    song_text text not null,
    release_date timestamp not null,
    link text not null,
    primary key (song_group, song)
    );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop schema library_schema cascade;
-- +goose StatementEnd
