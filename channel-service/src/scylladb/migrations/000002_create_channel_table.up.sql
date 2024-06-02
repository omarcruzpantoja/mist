CREATE TABLE IF NOT EXISTS channel (
    -- Need to refactor this so the primary key is better suited for partitions
    serverid UUID,
    id UUID,
    name ascii,
    created_at timestamp,
    updated_at timestamp,
    deleted_at timestamp,
    PRIMARY KEY (serverid, id)
);