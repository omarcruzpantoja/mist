CREATE TABLE IF NOT EXISTS server (
    -- Need to refactor this so the primary key is better suited for partitions
    id UUID,
    name ascii,
    PRIMARY KEY (id, name)
);