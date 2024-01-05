CREATE TABLE
    articles (
        id INTEGER PRIMARY KEY,
        title text NOT NULL,
        body text NOT NULL,
        created_at INTEGER DEFAULT (strftime('%s', 'now')),
        updated_at INTEGER DEFAULT (strftime('%s', 'now'))
    );