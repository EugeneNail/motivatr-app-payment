CREATE TABLE payments
(
    id          SERIAL PRIMARY KEY,
    date        DATE         NOT NULL,
    description VARCHAR(255) NOT NULL,
    category    SMALLINT     NOT NULL,
    value       FLOAT        NOT NULL,
    user_id     BIGINT       NOT NULL
);