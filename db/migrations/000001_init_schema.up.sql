USE jedel_komek;

CREATE TABLE users
(
    id           SERIAL PRIMARY KEY,
    fio          VARCHAR(255)        NOT NULL,
    phone_number VARCHAR(255) UNIQUE NOT NULL,
    password     TEXT                NOT NULL,
    created_at   TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE incident_reports
(
    id          SERIAL PRIMARY KEY,
    user_id     INTEGER REFERENCES users (id),
    description TEXT,
    media_url   TEXT,
    type_id     VARCHAR(255) REFERENCES incident_types (id),
    latitude    TEXT,
    longitude   TEXT,
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE incident_types
(
    id   SERIAL PRIMARY KEY,
    name VARCHAR(255)
);


CREATE TABLE emergency_calls
(
    id         SERIAL PRIMARY KEY,
    user_id    INTEGER REFERENCES users (id),
    latitude   TEXT,
    longitude  TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE education_contents
(
    id         SERIAL PRIMARY KEY,
    title      TEXT,
    body       TEXT,
    media_url  TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE police_department
(
    id           SERIAL PRIMARY KEY,
    name         VARCHAR(255),
    phone_number VARCHAR(255),
    address      VARCHAR(255),
    work_days    VARCHAR(255),
    work_time    VARCHAR(255),
    latitude     TEXT,
    longitude    TEXT,
    created_at   TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE news
(
    id         SERIAL PRIMARY KEY,
    title      TEXT      NOT NULL,
    content    TEXT      NOT NULL,
    media_url  TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE messages
(
    id          SERIAL PRIMARY KEY,
    sender_id   INTEGER   NOT NULL REFERENCES users (id),
    receiver_id INTEGER   NOT NULL REFERENCES users (id),
    text        TEXT      NOT NULL,
    created_at  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE notify_tokens
(
    id      SERIAL PRIMARY KEY,
    user_id INT REFERENCES users (id),
    token   VARCHAR(255)
);

CREATE TABLE notify_history
(
    id       SERIAL PRIMARY KEY,
    user_id  INT REFERENCES users (id),
    title    VARCHAR(255),
    body     VARCHAR(255),
    sender   INT,
    receiver INT
);