CREATE TABLE
    articles (
        id INTEGER PRIMARY KEY,
        title text NOT NULL,
        body text NOT NULL,
        created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
        updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
    );