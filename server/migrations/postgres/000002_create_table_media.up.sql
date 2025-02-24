CREATE TABLE media (
    media_id UUID PRIMARY KEY,
    resource_id UUID,
    created_at TIMESTAMP,
    modified_at TIMESTAMP,
    FOREIGN KEY(resource_id) REFERENCES resource(resource_id)
)