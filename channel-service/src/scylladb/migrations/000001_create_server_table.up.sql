CREATE TABLE IF NOT EXISTS server (
    -- Need to refactor this so the primary key is better suited for partitions
    id UUID,
    name ascii,
    created_at timestamp,
    updated_at timestamp,
    PRIMARY KEY (id)
);