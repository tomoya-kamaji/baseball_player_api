-- name: GetBook :one
SELECT
   *
FROM
   books
WHERE
   id = ?
LIMIT
   1;