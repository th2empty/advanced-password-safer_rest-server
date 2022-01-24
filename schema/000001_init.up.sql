CREATE TABLE users
(
    id serial not null unique,
    username VARCHAR(255) not null,
    password_hash VARCHAR(255) not null
);

CREATE TABLE password_list
(
    id serial not null unique,
    title VARCHAR(255) not null,
    description VARCHAR(255) not null
);

CREATE TABLE password_item
(
    id serial not null unique,
    title VARCHAR(255) not null,
    web_site VARCHAR(255) not null,
    login VARCHAR(255) not null,
    phone VARCHAR(255) not null,
    email VARCHAR(255) not null,
    pass VARCHAR(255) not null,
    secret_word VARCHAR(255) not null,
    backup_codes VARCHAR(255) not null
);

CREATE TABLE users_list
(
    id      serial                                           not null unique,
    user_id int references users (id) on delete cascade      not null,
    list_id int references password_list (id) on delete cascade not null
);

CREATE TABLE list_items
(
    id      serial                                           not null unique,
    item_id int references password_item (id) on delete cascade not null,
    list_id int references password_list (id) on delete cascade not null
);