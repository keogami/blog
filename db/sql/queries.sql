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
SELECT * FROM metas
WHERE post_id NOT IN (
  SELECT post_id FROM deleted
);

-- name: GetPostBySlug :one
SELECT * FROM posts, metas
WHERE
  id = post_id AND
  slug = $1 AND
  NOT EXISTS (SELECT post_id FROM deleted WHERE post_id = id);

-- name: UpdatePost :one
UPDATE posts
SET content = $2
FROM metas
WHERE metas.slug = $1
RETURNING id;

-- name: UpdateMeta :exec
UPDATE metas
SET (
  slug, title, summary, group_name, tags
) = (
  $2, $3, $4, $5, $6
)
WHERE post_id = $1;

-- name: MarkDeleted :exec
INSERT INTO deleted 
SELECT post_id FROM metas 
WHERE slug = $1;
