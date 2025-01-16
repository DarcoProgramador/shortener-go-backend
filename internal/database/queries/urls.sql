-- name: GetURLByShortCode :one
SELECT 
    id,
    url,
    shortCode,
    createdAt,
    updatedAt
FROM urls
WHERE shortCode = ?;

-- name: CreateURL :one
INSERT INTO urls (url, shortCode)
VALUES (?, ?)
RETURNING id, url, shortCode, createdAt, updatedAt;

-- name: UpdateURLByShortCode :one
UPDATE urls
SET url = ?, updatedAt = ?
WHERE shortCode = ?
RETURNING id, url, shortCode, createdAt, updatedAt;

-- name: IncrementURLAccessCountByShortCode :exec
UPDATE urls
SET accessCount = accessCount + 1
WHERE shortCode = ?;

-- name: DeleteURLByShortCode :exec
DELETE FROM urls
WHERE shortCode = ?;

-- name: GetURLStatsByShortCode :one
SELECT 
    id,
    url,
    shortCode,
    createdAt,
    updatedAt,
    accessCount
FROM urls
WHERE shortCode = ?;