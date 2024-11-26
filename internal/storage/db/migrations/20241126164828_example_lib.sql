-- +goose Up
-- +goose StatementBegin
insert into library_schema.library(song_group, song, song_text, release_date, link)
values ('group1', 'song1','some\ntext\separated\nby\nseparator', '2020.11.13', 'someLing1'),
       ('group2', 'song2','some\ntext\separated\nby\nseparator', '2020.11.14', 'someLing2'),
       ('group3', 'song3','some\ntext\separated\nby\nseparator', '2020.11.15', 'someLink3'),
       ('group4', 'song4','some\ntext\separated\nby\nseparator', '2020.10.13', 'someLing4'),
       ('group5', 'song5','some\ntext\separated\nby\nseparator', '2020.09.13', 'someLing5'),
       ('group6', 'song6','some\ntext\separated\nby\nseparator', '2020.11.17', 'someLing6'),
       ('group7', 'song7','some\ntext\separated\nby\nseparator', '2020.12.13', 'someLing7'),
       ('group8', 'song8','some\ntext\separated\nby\nseparator', '2020.10.10', 'someLing8');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
truncate table library_schema.library cascade;
-- +goose StatementEnd
