// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0
// source: query.sql

package store

import (
	"context"
	"database/sql"
)

const getSubmissionByDocumentID = `-- name: GetSubmissionByDocumentID :one

SELECT inv_id, doc_uid, sub_uid, accepted, status_, long_id, upsert_at FROM submissions WHERE doc_uid = ?
`

// -- name: GetSubmissionBySubmissionID :one
// SELECT * FROM submissions WHERE sub_uid = ?; -- not indexed
func (q *Queries) GetSubmissionByDocumentID(ctx context.Context, docUid string) (Submission, error) {
	row := q.db.QueryRowContext(ctx, getSubmissionByDocumentID, docUid)
	var i Submission
	err := row.Scan(
		&i.InvID,
		&i.DocUid,
		&i.SubUid,
		&i.Accepted,
		&i.Status,
		&i.LongID,
		&i.UpsertAt,
	)
	return i, err
}

const getSubmissionByInvoiceID = `-- name: GetSubmissionByInvoiceID :one
SELECT inv_id, doc_uid, sub_uid, accepted, status_, long_id, upsert_at FROM submissions WHERE inv_id = ?
`

func (q *Queries) GetSubmissionByInvoiceID(ctx context.Context, invID string) (Submission, error) {
	row := q.db.QueryRowContext(ctx, getSubmissionByInvoiceID, invID)
	var i Submission
	err := row.Scan(
		&i.InvID,
		&i.DocUid,
		&i.SubUid,
		&i.Accepted,
		&i.Status,
		&i.LongID,
		&i.UpsertAt,
	)
	return i, err
}

const getSubmissionByStatusNull = `-- name: GetSubmissionByStatusNull :many
SELECT inv_id, doc_uid, sub_uid, accepted, status_, long_id, upsert_at FROM submissions WHERE status_ IS NULL
`

func (q *Queries) GetSubmissionByStatusNull(ctx context.Context) ([]Submission, error) {
	rows, err := q.db.QueryContext(ctx, getSubmissionByStatusNull)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Submission
	for rows.Next() {
		var i Submission
		if err := rows.Scan(
			&i.InvID,
			&i.DocUid,
			&i.SubUid,
			&i.Accepted,
			&i.Status,
			&i.LongID,
			&i.UpsertAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getSubmissions = `-- name: GetSubmissions :many
SELECT inv_id, doc_uid, sub_uid, accepted, status_, long_id, upsert_at FROM submissions ORDER BY upsert_at DESC LIMIT 100
`

func (q *Queries) GetSubmissions(ctx context.Context) ([]Submission, error) {
	rows, err := q.db.QueryContext(ctx, getSubmissions)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Submission
	for rows.Next() {
		var i Submission
		if err := rows.Scan(
			&i.InvID,
			&i.DocUid,
			&i.SubUid,
			&i.Accepted,
			&i.Status,
			&i.LongID,
			&i.UpsertAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const insertSubmission = `-- name: InsertSubmission :one
INSERT INTO submissions ( inv_id, doc_uid, sub_uid, accepted )
VALUES ( ?, ?, ?, ? ) RETURNING doc_uid
`

type InsertSubmissionParams struct {
	InvID    string
	DocUid   string
	SubUid   string
	Accepted sql.NullInt64
}

func (q *Queries) InsertSubmission(ctx context.Context, arg InsertSubmissionParams) (string, error) {
	row := q.db.QueryRowContext(ctx, insertSubmission,
		arg.InvID,
		arg.DocUid,
		arg.SubUid,
		arg.Accepted,
	)
	var doc_uid string
	err := row.Scan(&doc_uid)
	return doc_uid, err
}

const invoiceExistsByID = `-- name: InvoiceExistsByID :one
SELECT EXISTS( SELECT 1 FROM submissions WHERE inv_id = ?)
`

func (q *Queries) InvoiceExistsByID(ctx context.Context, invID string) (int64, error) {
	row := q.db.QueryRowContext(ctx, invoiceExistsByID, invID)
	var column_1 int64
	err := row.Scan(&column_1)
	return column_1, err
}

const updateSubmissionStatusLongID = `-- name: UpdateSubmissionStatusLongID :exec
UPDATE submissions SET status_ = ?, long_id = ?, upsert_at = CURRENT_TIMESTAMP WHERE doc_uid = ?
`

type UpdateSubmissionStatusLongIDParams struct {
	Status sql.NullString
	LongID sql.NullString
	DocUid string
}

func (q *Queries) UpdateSubmissionStatusLongID(ctx context.Context, arg UpdateSubmissionStatusLongIDParams) error {
	_, err := q.db.ExecContext(ctx, updateSubmissionStatusLongID, arg.Status, arg.LongID, arg.DocUid)
	return err
}
