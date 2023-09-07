CREATE TABLE api_keys
(
    id              serial primary key,
    hash            varchar(255) not null unique,
    created_at      timestamp             default CURRENT_TIMESTAMP,
    expire_at       timestamp,
    updated_at      timestamp             default CURRENT_TIMESTAMP,
    is_active       boolean      not null default true,
    hash_version_id int          not null
);