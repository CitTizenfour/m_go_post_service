CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

create table if not exists post 
(
    "guid"               uuid      default uuid_generate_v4() primary key, 
    "id"                 serial,
    "text"               text,
    "content"            text,
    "publication_date"   text,
    "user_id"            int,
    "created_at"         timestamp default CURRENT_TIMESTAMP,
    "updated_at"         timestamp default CURRENT_TIMESTAMP
)