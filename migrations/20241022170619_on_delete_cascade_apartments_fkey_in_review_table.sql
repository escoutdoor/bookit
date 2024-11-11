-- +goose Up
-- +goose StatementBegin
ALTER TABLE REVIEWS
DROP CONSTRAINT reviews_apartment_id_fkey,
ADD CONSTRAINT reviews_apartment_id_fkey
FOREIGN KEY (apartment_id) REFERENCES APARTMENTS(id) ON DELETE CASCADE;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE REVIEWS
DROP CONSTRAINT reviews_apartment_id_fkey,
ADD CONSTRAINT reviews_apartment_id_fkey
FOREIGN KEY (apartment_id) REFERENCES APARTMENTS(id);
-- +goose StatementEnd
