-- +goose Up
-- +goose StatementBegin
CREATE TABLE products (
id serial NOT NULL PRIMARY KEY,
title text NOT NULL,
description text NOT NULL,
price int NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE products;
-- +goose StatementEnd
