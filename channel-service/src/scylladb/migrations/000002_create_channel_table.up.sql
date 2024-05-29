CREATE TABLE IF NOT EXISTS channel (
    -- Need to refactor this so the primary key is better suited for partitions
    serverid UUID,
    id UUID,
    name ascii,
    PRIMARY KEY (serverid, id)
);