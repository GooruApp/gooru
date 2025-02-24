CREATE TABLE media_tag (
    media_id UUID,
    tag_id UUID,
    created_at TIMESTAMP,
    modified_at TIMESTAMP,
    FOREIGN KEY(media_id) REFERENCES media(media_id),
    FOREIGN KEY(tag_id) REFERENCES tag(tag_id),
    PRIMARY KEY (media_id, tag_id)
)