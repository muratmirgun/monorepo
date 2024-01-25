-- +goose Up
CREATE TABLE IF NOT EXISTS customer (
       id               UUID PRIMARY KEY DEFAULT (gen_random_uuid()),
       email           varchar NOT NULL,
       password        varchar NOT NULL,
       created_at       TIMESTAMP NOT NULL DEFAULT NOW(),
       updated_at       TIMESTAMP NOT NULL DEFAULT NOW(),
       deleted_at       TIMESTAMP
);

CREATE TABLE IF NOT EXISTS customer_user (
       id               UUID PRIMARY KEY DEFAULT (gen_random_uuid()),
       customer_id      UUID NOT NULL,
       user_id          UUID NOT NULL,
       created_at       TIMESTAMP NOT NULL DEFAULT NOW(),
       updated_at       TIMESTAMP NOT NULL DEFAULT NOW(),
       deleted_at       TIMESTAMP,
       CONSTRAINT customer_user_key FOREIGN KEY (customer_id) REFERENCES customer(id)
);

