CREATE TABLE urls (
    id bigserial not null primary key,
    long_url varchar not null,
    short_url varchar not null unique,
    delete_key varchar not null
);