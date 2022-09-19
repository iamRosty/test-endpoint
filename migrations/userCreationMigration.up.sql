CREATE TABLE users (
    id integer not null primary key,
    name varchar not null,
    surname varchar not null,
    isadmin boolean not null,
    email varchar not null,
    password varchar not null
);