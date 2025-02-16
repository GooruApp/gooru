CREATE TABLE media (
    media_id uuid PRIMARY KEY,
    resource_id uuid,
    FOREIGN KEY(resource_id) REFERENCES resource(resource_id)
)