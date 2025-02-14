-- name: CreateAccount :one
INSERT INTO ACCOUNTS (
  USER_ID,
  CATEGORY_ID,
  TITLE,
  TYPE,
  DESCRIPTION,
  VALUE,
  DATE
) VALUES (
  $1,
  $2,
  $3,
  $4,
  $5,
  $6,
  $7
) RETURNING *;

-- name: GetAccount :one
SELECT
  *
FROM
  ACCOUNTS
WHERE
  ID = $1 LIMIT 1;

-- name: GetAccounts :many
SELECT
  A.ID,
  A.USER_ID,
  A.TITLE,
  A.TYPE,
  A.DESCRIPTION,
  A.VALUE,
  A.DATE,
  A.CREATED_AT,
  C.TITLE AS CATEGORY_TITLE
FROM
  ACCOUNTS A
  LEFT JOIN CATEGORIES C
  ON C.ID = A.CATEGORY_ID
WHERE
  A.USER_ID = @USER_ID
  AND A.TYPE = @TYPE
  AND LOWER(A.TITLE) LIKE CONCAT('%',
  LOWER(@TITLE::TEXT),
  '%')
  AND LOWER(A.DESCRIPTION) LIKE CONCAT('%',
  LOWER(@DESCRIPTION::TEXT),
  '%')
  AND A.CATEGORY_ID = COALESCE(SQLC.NARG('category_id'),
  A.CATEGORY_ID)
  AND A.DATE = COALESCE(SQLC.NARG('date'),
  A.DATE);

-- name: GetAccountsByUserIdAndType :many
SELECT
  A.ID,
  A.USER_ID,
  A.TITLE,
  A.TYPE,
  A.DESCRIPTION,
  A.VALUE,
  A.DATE,
  A.CREATED_AT,
  C.TITLE AS CATEGORY_TITLE
FROM
  ACCOUNTS   A
  LEFT JOIN CATEGORIES C
  ON C.ID = A.CATEGORY_ID
WHERE
  A.USER_ID = $1
  AND A.TYPE = $2;

-- name: GetAccountsByUserIdAndTypeAndCategoryId :many
SELECT
  A.ID,
  A.USER_ID,
  A.TITLE,
  A.TYPE,
  A.DESCRIPTION,
  A.VALUE,
  A.DATE,
  A.CREATED_AT,
  C.TITLE AS CATEGORY_TITLE
FROM
  ACCOUNTS   A
  LEFT JOIN CATEGORIES C
  ON C.ID = A.CATEGORY_ID
WHERE
  A.USER_ID = $1
  AND A.TYPE = $2
  AND A.CATEGORY_ID = $3;

-- name: GetAccountsByUserIdAndTypeAndCategoryIdAndTitle :many
SELECT
  A.ID,
  A.USER_ID,
  A.TITLE,
  A.TYPE,
  A.DESCRIPTION,
  A.VALUE,
  A.DATE,
  A.CREATED_AT,
  C.TITLE AS CATEGORY_TITLE
FROM
  ACCOUNTS A
  LEFT JOIN CATEGORIES C
  ON C.ID = A.CATEGORY_ID
WHERE
  A.USER_ID = $1
  AND A.TYPE = $2
  AND A.CATEGORY_ID = $3
  AND LOWER(A.TITLE) LIKE CONCAT('%',
  LOWER($4::TEXT),
  '%');

-- name: GetAccountsByUserIdAndTypeAndAndCategoryIdAndTitleAndDescription :many
SELECT
  A.ID,
  A.USER_ID,
  A.TITLE,
  A.TYPE,
  A.DESCRIPTION,
  A.VALUE,
  A.DATE,
  A.CREATED_AT,
  C.TITLE AS CATEGORY_TITLE
FROM
  ACCOUNTS A
  LEFT JOIN CATEGORIES C
  ON C.ID = A.CATEGORY_ID
WHERE
  A.USER_ID = $1
  AND A.TYPE = $2
  AND A.CATEGORY_ID = $3
  AND LOWER(A.TITLE) LIKE CONCAT('%',
  LOWER($4::TEXT),
  '%')
  AND LOWER(A.DESCRIPTION) LIKE CONCAT('%',
  LOWER($5::TEXT),
  '%');

-- name: GetAccountsByUserIdAndTypeAndTitle :many
SELECT
  A.ID,
  A.USER_ID,
  A.TITLE,
  A.TYPE,
  A.DESCRIPTION,
  A.VALUE,
  A.DATE,
  A.CREATED_AT,
  C.TITLE AS CATEGORY_TITLE
FROM
  ACCOUNTS A
  LEFT JOIN CATEGORIES C
  ON C.ID = A.CATEGORY_ID
WHERE
  A.USER_ID = $1
  AND A.TYPE = $2
  AND LOWER(A.TITLE) LIKE CONCAT('%',
  LOWER($2::TEXT),
  '%');

-- name: GetAccountsByUserIdAndTypeAndDescription :many
SELECT
  A.ID,
  A.USER_ID,
  A.TITLE,
  A.TYPE,
  A.DESCRIPTION,
  A.VALUE,
  A.DATE,
  A.CREATED_AT,
  C.TITLE AS CATEGORY_TITLE
FROM
  ACCOUNTS A
  LEFT JOIN CATEGORIES C
  ON C.ID = A.CATEGORY_ID
WHERE
  A.USER_ID = $1
  AND A.TYPE = $2
  AND LOWER(A.DESCRIPTION) LIKE CONCAT('%',
  LOWER($3::TEXT),
  '%');

-- name: GetAccountsByUserIdAndTypeAndTitleAndDate :many
SELECT
  A.ID,
  A.USER_ID,
  A.TITLE,
  A.TYPE,
  A.DESCRIPTION,
  A.VALUE,
  A.DATE,
  A.CREATED_AT,
  C.TITLE AS CATEGORY_TITLE
FROM
  ACCOUNTS A
  LEFT JOIN CATEGORIES C
  ON C.ID = A.CATEGORY_ID
WHERE
  A.USER_ID = $1
  AND A.TYPE = $2
  AND LOWER(A.TITLE) LIKE CONCAT('%',
  LOWER($3::TEXT),
  '%')
  AND A.DATE = $4;

-- name: GetAccountsReports :one
SELECT
  SUM(VALUE) AS SUM_VALUE
FROM
  ACCOUNTS
WHERE
  USER_ID = $1
  AND TYPE = $2;

-- name: GetAccountsGraph :one
SELECT
  COUNT(*)
FROM
  ACCOUNTS
WHERE
  USER_ID = $1
  AND TYPE = $2;

-- name: UpdateAccount :one
UPDATE ACCOUNTS
SET
  TITLE = $2,
  DESCRIPTION = $3,
  VALUE = $4
WHERE
  ID = $1 RETURNING *;

-- name: DeleteAccount :exec
DELETE FROM ACCOUNTS
WHERE
  ID = $1;

-- docker run --rm -v "C:\Users\peves\Documents\go_finance:/src" -w /src kjconroy/sqlc generate