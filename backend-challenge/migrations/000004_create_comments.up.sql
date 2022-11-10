CREATE TABLE IF NOT EXISTS comments (
    id bigserial not null primary key,
    author varchar not null,
    author_img varchar,
    date varchar not null,
    comment varchar,
    comment_img varchar
);