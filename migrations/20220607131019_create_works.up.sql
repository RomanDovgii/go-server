CREATE TABLE works (
  id BIGSERIAL NOT NULL PRIMARY KEY,
  creator_id INTEGER PREFERENCES users (id),
  name STRING,
  description TEXT,
  document_links []
);