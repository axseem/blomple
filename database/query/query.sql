-- name: GetArticle :one
SELECT * FROM articles
WHERE id = ? LIMIT 1;

-- name: ListArticles :many
SELECT * FROM articles
ORDER BY created_at DESC;

-- name: CreateArticle :exec
INSERT INTO articles (
  title, body
) VALUES (
  ?, ?
);

-- name: UpdateArticle :exec
UPDATE articles
set title = ?, body = ?
WHERE id = ?;

-- name: DeleteArticle :exec
DELETE FROM articles
WHERE id = ?;