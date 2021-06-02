CREATE TABLE IF NOT EXISTS url_entries (
    short_url text NOT NULL PRIMARY KEY,
    original_url text NOT NULL
);
