create table if not exists users (
    id bigserial primary key,
    first_name text not null,
    last_name text not null,
    phone text not null unique,
    email text not null unique,
    password_hash text not null,
    created_at timestamp not null default now()
);

create type if not exists sex AS enum ('Other', 'Male', 'Female');

create table if not exists profile (
    id bigserial primary key,
    sex sex not null,
    date_of_birth date not null,
    shoes_size integer not null,
    user_id integer not null,
    check (shoes_size > 0),
    foreign key (user_id) references users(id)
);

create table if not exists address (
    id bigserial primary key,
    country text not null,
    city text not null,
    street text not null,
    house text not null,
    appartment text,
    user_id integer not null,
    foreign key (user_id) references users(id),
);

create index if not exists idx_users_signin on users using btree (phone, password);

create index if not exists idx_profile_user-id on profile using btree (user_id);

create index if not exists idx_address_user-id on address using btree (user_id);