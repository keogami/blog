-- +goose Up
CREATE TABLE IF NOT EXISTS posts (
  id BIGSERIAL PRIMARY KEY,
  content text NOT NULL
);

CREATE TABLE IF NOT EXISTS metas (
  post_id BIGSERIAL NOT NULL REFERENCES posts(id),
  slug text NOT NULL,
  title VARCHAR(127) NOT NULL,
  summary VARCHAR(255) NOT NULL,
  group_name VARCHAR(16),
  tags VARCHAR(16)[],
  post_time TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

