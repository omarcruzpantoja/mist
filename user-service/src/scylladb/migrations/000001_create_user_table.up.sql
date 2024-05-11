CREATE TABLE IF NOT EXISTS user (
    -- Need to refactor this so the primary key is better suited for partitions
    id UUID,
    username ascii,
    password ascii,
    PRIMARY KEY (id, username)
);