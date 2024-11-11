-- +goose Up
-- +goose StatementBegin
CREATE TYPE user_role as ENUM('admin', 'host', 'guest');

CREATE TABLE IF NOT EXISTS USERS(
    id uuid primary key default gen_random_uuid(),
    email varchar(255) not null unique,
    password varchar(255) not null,
    role user_role not null default 'guest',
    first_name varchar(255) not null,
    last_name varchar(255) not null,
    date_of_birth DATE,
    avatar_url varchar(255),
    phone_number varchar(20) unique,
    created_at timestamp not null default now()
);

CREATE TABLE IF NOT EXISTS CATEGORIES(
    id uuid primary key default gen_random_uuid(),
    name varchar(255) not null,
    created_at timestamp not null default now()
);

CREATE TABLE IF NOT EXISTS APARTMENTS(
    id uuid primary key default gen_random_uuid(),
    name varchar(255) not null,
    description text,
    beds integer,
    bedrooms integer,
    bathrooms integer,
    max_guests integer not null,
    rental_price decimal(10, 2) not null,
    latitude decimal(9, 6) not null,
    longitude decimal(9, 6) not null,
    host_id uuid not null,
    category_id uuid,
    created_at timestamp not null default now(),
    FOREIGN KEY(host_id) REFERENCES USERS(id),
    FOREIGN KEY(category_id) REFERENCES CATEGORIES(id) ON DELETE SET NULL
);

CREATE TABLE IF NOT EXISTS RESERVATIONS(
    id uuid primary key default gen_random_uuid(),
    renter_id uuid not null,
    apartment_id uuid not null,
    start_date timestamp not null,
    end_date timestamp not null,
    guests integer,
    total decimal(10, 2) not null,
    created_at timestamp not null default now(),
    FOREIGN KEY(renter_id) REFERENCES USERS(id),
    FOREIGN KEY(apartment_id) REFERENCES APARTMENTS(id)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE RESERVATIONS;
DROP TABLE APARTMENTS;
DROP TABLE CATEGORIES;

DROP TABLE USERS;
DROP TYPE USER_ROLE;
-- +goose StatementEnd
