-- +goose Up

CREATE TABLE users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    email TEXT UNIQUE NOT NULL,
    avatar TEXT,
    refresh_token TEXT,
    role INTEGER NOT NULL DEFAULT 0 CHECK (role IN (0, 1)),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE api_categories (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE apis (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    version TEXT,
    description TEXT,
    documentation_link TEXT,
    forum_reference TEXT,
    tags TEXT,
    swagger TEXT,
    apm_link TEXT,
    team      TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE api_category_mappings (
    api_id INTEGER,
    category_id INTEGER,
    PRIMARY KEY (api_id, category_id),
    FOREIGN KEY (api_id) REFERENCES apis(id) ON DELETE CASCADE,
    FOREIGN KEY (category_id) REFERENCES api_categories(id) ON DELETE CASCADE
);

CREATE INDEX idx_api_category_mappings_api_id ON api_category_mappings(api_id);
CREATE INDEX idx_api_category_mappings_category_id ON api_category_mappings(category_id);

-- +goose Down

DROP INDEX IF EXISTS idx_api_category_mappings_category_id;
DROP INDEX IF EXISTS idx_api_category_mappings_api_id;
DROP TABLE IF EXISTS api_category_mappings;
DROP TABLE IF EXISTS apis;
DROP TABLE IF EXISTS api_categories;
DROP TABLE IF EXISTS users;