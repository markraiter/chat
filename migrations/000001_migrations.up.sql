CREATE TABLE users (
    id            INT PRIMARY KEY AUTO_INCREMENT,
    avatar        VARCHAR(255),
    username      VARCHAR(255) NOT NULL UNIQUE,
    password      VARCHAR(255) NOT NULL
);

CREATE TABLE messages (
    id            INT PRIMARY KEY AUTO_INCREMENT,
    user_id       INT NOT NULL,
    username      VARCHAR(255) UNIQUE,
    body          TEXT,
    FOREIGN KEY (user_id) REFERENCES users(id)
);