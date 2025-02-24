CREATE TABLE media (
    media_id TEXT PRIMARY KEY,
    resource_id TEXT,
    created_at TEXT,
    modified_at TEXT,
    FOREIGN KEY(resource_id) REFERENCES resource(resource_id)
)