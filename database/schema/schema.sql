CREATE TABLE
    article (
        id INTEGER PRIMARY KEY,
        title text NOT NULL,
        content text NOT NULL,
        created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
        updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
    );