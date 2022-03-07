-- name: CreatePost :one
INSERT INTO posts (content) VALUES ($1)
RETURNING id;

-- name: CreateMeta :one
INSERT INTO metas (
  post_id, slug, title, summary, group_name, tags
) VALUES (
  $1, $2, $3, $4, $5, $6
)
RETURNING *;

-- name: ListMetas :many
SELECT * FROM metas;

-- name: GetPostBySlug :one
SELECT * FROM posts, metas
WHERE
  id = post_id AND
  slug = $1
;

