-- name: GetSubmissions :many
SELECT * FROM submissions ORDER BY upsert_at DESC LIMIT 100;

-- name: GetSubmissionByInvoiceID :one
SELECT * FROM submissions WHERE inv_id = ?;

-- name: InvoiceExistsByID :one
SELECT EXISTS( SELECT 1 FROM submissions WHERE inv_id = ?);

-- -- name: GetSubmissionBySubmissionID :one
-- SELECT * FROM submissions WHERE sub_uid = ?; -- not indexed

-- name: GetSubmissionByDocumentID :one
SELECT * FROM submissions WHERE doc_uid = ?;

-- name: GetSubmissionByStatusNull :many
SELECT * FROM submissions WHERE status_ IS NULL;

-- name: InsertSubmission :one
INSERT INTO submissions ( inv_id, doc_uid, sub_uid, accepted )
VALUES ( ?, ?, ?, ? ) RETURNING doc_uid;

-- name: UpdateSubmissionStatusLongID :exec
UPDATE submissions SET status_ = ?, long_id = ?, upsert_at = CURRENT_TIMESTAMP WHERE doc_uid = ?;