CREATE TABLE products
(
    id          serial       not null unique,
    name        varchar(255) not null unique,
    description varchar(255),
    price       bigint,
    category_id int,
    created_at  timestamp,
    modified_at timestamp,
    deleted_at  timestamp
);