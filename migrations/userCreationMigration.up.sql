CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name varchar not null,
    surname varchar not null,
    isadmin boolean not null,
    email varchar not null,
    password varchar not null
);