-- +migrate Up
CREATE TABLE users (
    id varchar(36) NOT NULL PRIMARY KEY,
    username varchar(50) NOT NULL,
    password varchar(255) NOT NULL,
    full_name varchar(255) NULL,
    is_active bool NOT NULL DEFAULT FALSE,
    refresh_token varchar(255) NULL,
    created_at timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT users_unique UNIQUE (username)
);

INSERT INTO users (id, username, password, full_name, is_active) VALUES (uuid(),'admin', '$2a$04$LiSuUvol8QlO76ePndhH5OzSc6vdpbewyQbFWSdyHgn3q3xTfXhnG', 'Admin', TRUE);

-- +migrate Down
DROP TABLE users;

