CREATE TABLE IF NOT EXISTS channel (
    -- Need to refactor this so the primary key is better suited for partitions
    mserverid UUID,
    id UUID,
    name ascii,
    PRIMARY KEY (mserverid, id, name)
);