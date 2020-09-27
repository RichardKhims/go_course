CREATE TABLE IF NOT EXISTS currency (
    id BIGSERIAL PRIMARY KEY,
    symbol CHAR(3) UNIQUE NOT NULL,
    curname VARCHAR(32)
);

CREATE TABLE IF NOT EXISTS course (
    id BIGSERIAL PRIMARY KEY,
    cur1 CHAR(3) NOT NULL,
    cur2 CHAR(3) NOT NULL,
    mean DOUBLE NOT NULL
);
