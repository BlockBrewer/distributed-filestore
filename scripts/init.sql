CREATE TABLE IF NOT EXISTS file_chunks ( id SERIAL PRIMARY KEY,
                                                           file_id TEXT, data BYTEA,
                                                                              part INT);


CREATE TABLE IF NOT EXISTS file_metadata ( id TEXT PRIMARY KEY,
                                                           name TEXT, size BIGINT);


CREATE INDEX IF NOT EXISTS idx_file_chunks_file_id ON file_chunks(file_id);


CREATE INDEX IF NOT EXISTS idx_file_chunks_part ON file_chunks(part);