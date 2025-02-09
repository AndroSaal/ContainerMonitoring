CREATE TABLE IF NOT EXISTS container_ip (
    id SERIAL PRIMARY KEY,
    ip VARCHAR(255) NOT NULL UNIQUE,
    last_success TIMESTAMP
);

CREATE TABLE IF NOT EXISTS container_ping (
    id SERIAL PRIMARY KEY,
    ip VARCHAR(255),
    ping_time TIMESTAMP,
    status VARCHAR(255), 
    FOREIGN KEY (ip) REFERENCES container_ip(ip) ON DELETE CASCADE
);