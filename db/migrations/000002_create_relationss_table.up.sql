CREATE TABLE relation_type
(
    id             serial primary key,
    name           text not null,
    name_canonical text not null unique,
    comment        text
);