-- name: GetBook :many
SELECT
   *
FROM
   books
WHERE
   id IN (sqlc.slice("bookIds"));

-- name: GetBookByTitle :many
SELECT
   *
FROM
   books
WHERE
   title = sqlc.arg(title);