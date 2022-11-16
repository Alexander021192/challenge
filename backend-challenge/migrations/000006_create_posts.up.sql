CREATE TABLE IF NOT EXISTS posts (
    id bigserial not null primary key,
    author varchar not null,
    title varchar not null,
    location varchar,
    post_text varchar not null,
    post_img varchar
);