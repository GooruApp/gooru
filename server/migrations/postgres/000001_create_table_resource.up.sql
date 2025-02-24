CREATE TABLE resource (
    resource_id UUID PRIMARY KEY,
    uri TEXT,
    bytes INTEGER,
    created_at TIMESTAMP,
    modified_at TIMESTAMP
)