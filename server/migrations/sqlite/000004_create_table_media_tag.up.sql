CREATE TABLE media_tag (
    media_id TEXT,
    tag_id TEXT,
    FOREIGN KEY(media_id) REFERENCES media(media_id),
    FOREIGN KEY(tag_id) REFERENCES tag(tag_id),
    PRIMARY KEY (media_id, tag_id)
)