-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS REVIEWS(
    id uuid primary key default gen_random_uuid(),
    content varchar(255) not null,
    rating int not null,
    renter_id uuid not null,
    apartment_id uuid not null,
    created_at timestamp not null default now(),
    FOREIGN KEY(apartment_id) REFERENCES APARTMENTS(id),
    FOREIGN KEY(renter_id) REFERENCES USERS(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE REVIEWS;
-- +goose StatementEnd
