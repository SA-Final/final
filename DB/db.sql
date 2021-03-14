
CREATE TABLE if not exists books
(
    id serial primary key,
    name varchar(100) NOT NULL,
    author varchar(100) NOT NULL
);

CREATE TABLE if not exists users
(
    id serial primary key,
    email varchar(100) NOT NULL,
    username varchar(100) NOT NULL,
    password varchar(255) NOT NULL
);

CREATE TABLE if not exists users_books
(
    user_id integer NOT NULL,
    book_id integer NOT NULL
);

