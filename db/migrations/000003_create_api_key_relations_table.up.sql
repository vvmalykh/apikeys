CREATE TABLE api_key_relations
(
    id               serial primary key,
    api_key_id       int not null,
    relation_type_id int not null,
    related_id       int not null
);

create index api_key_relations_api_key_id__index
    on api_key_relations (api_key_id);