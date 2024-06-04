CREATE TABLE users (
    user_id     UUID        PRIMARY KEY,
    fullname    VARCHAR,
    username    VARCHAR     NOT NULL,
    gmail       VARCHAR     NOT NULL,
    password    VARCHAR     NOT NULL
);

CREATE TABLE todos (
    todo_id         UUID    PRIMARY KEY,
    task            VARCHAR     NOT NULL,
    created_at      TIMESTAMP   DEFAULT current_timestamp,
    is_completed    BOOLEAN     DEFAULT false
);