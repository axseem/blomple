-- name: GetArticle :one
SELECT * FROM article
WHERE id = ? LIMIT 1;

-- name: ListArticles :many
SELECT * FROM article
ORDER BY created_at DESC;

-- name: CreateArticle :exec
INSERT INTO article (
  title, content
) VALUES (
  ?, ?
);

-- name: UpdateArticle :exec
UPDATE article
set title = ?, content = ?
WHERE id = ?;

-- name: DeleteArticle :exec
DELETE FROM article
WHERE id = ?;