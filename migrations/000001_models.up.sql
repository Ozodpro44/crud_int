
CREATE TABLE users_todo (
    user_id     UUID        PRIMARY KEY,
    fullname    VARCHAR,
    username    VARCHAR     UNIQUE          NOT NULL,
    gmail       VARCHAR                     NOT NULL,
    password    VARCHAR                     NOT NULL
);

CREATE TABLE todos (
    user_id         UUID    REFERENCES users_todo(user_id),
    todo_id         UUID    PRIMARY KEY,
    task            VARCHAR     NOT NULL,
    created_at      TIMESTAMP   DEFAULT current_timestamp,
    is_completed    BOOLEAN     DEFAULT false
);

