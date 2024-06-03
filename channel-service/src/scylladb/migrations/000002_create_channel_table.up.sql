CREATE TABLE IF NOT EXISTS channel (
    -- Need to refactor this so the primary key is better suited for partitions
    server_id UUID,
    id UUID,
    name ascii,
    created_at timestamp,
    updated_at timestamp,
    PRIMARY KEY (server_id, id)
);