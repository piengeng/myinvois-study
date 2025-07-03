CREATE TABLE IF NOT EXISTS submissions (
  inv_id    text PRIMARY KEY,
  doc_uid   text NOT NULL,
  sub_uid   text NOT NULL,
  accepted  integer DEFAULT 0, -- 1: accepted, 0: rejected
  status_   text DEFAULT NULL, -- 'Valid|??'
  long_id   text DEFAULT NULL, -- validation link https://{{portal.url}}/{{doc_uid}}/share/{{long_id}}
  upsert_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_doc_uid ON submissions (doc_uid);
-- CREATE INDEX IF NOT EXISTS idx_sub_uid ON submissions (sub_uid); -- less usage
-- CREATE INDEX IF NOT EXISTS idx_accepted ON submissions (accepted); -- erm maybe use cpu
CREATE INDEX IF NOT EXISTS idx_status_ ON submissions (status_);
CREATE INDEX IF NOT EXISTS idx_upsert_at ON submissions (upsert_at);