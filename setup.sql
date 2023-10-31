CREATE TABLE users (
    id bigserial primary key,
    uuid varchar(64) not null unique,
    name varchar(255),
    email varchar(255) not null unique,
    password varchar(255) not null,
    created_at timestamptz not null default now(),
    updated_at timestamptz not null default now(),
    deleted_at timestamptz
)

CREATE TABLE sessions (
    id bigserial primary key,
    uuid varchar(64) not null unique,
    email varchar(255) not null,
    user_id bigint not null references users(id),
    created_at timestamptz not null default now(),
    deleted_at timestamptz
)

CREATE TABLE threads (
    id bigserial primary key,
    uuid varchar(64) not null unique,
    topic text,
    created_by bigint not null references users(id),
    created_at timestamptz not null default now(),
    updated_at timestamptz not null default now(),
    deleted_at timestamptz
)

CREATE TABLE posts (
    id bigserial primary key,
    uuid varchar(64) not null unique,
    body text,
    created_by bigint not null references users(id),
    created_at timestamptz not null default now(),
    updated_at timestamptz not null default now(),
    deleted_at timestamptz
)

