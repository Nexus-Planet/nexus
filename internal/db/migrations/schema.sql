-- Users table
CREATE TABLE
  users (
    id VARCHAR(36) PRIMARY KEY,
    email VARCHAR(255) NOT NULL UNIQUE,
    username VARCHAR(50) UNIQUE,
    password_hash TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
  );

-- Guilds table
CREATE TABLE
  guilds (
     id VARCHAR(36) PRIMARY KEY,
    guild_name VARCHAR(255) NOT NULL,
    invite_link VARCHAR(255)
  );

-- User Guilds joined table
CREATE TABLE
  user_guilds (
    user_id VARCHAR(36) NOT NULL,
    guild_id VARCHAR(36) NOT NULL,
    role VARCHAR(50) DEFAULT 'member',
    joined_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (user_id, guild_id),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (guild_id) REFERENCES guilds(id) ON DELETE CASCADE
  );

-- Messages table
CREATE TABLE
  messages (
    id VARCHAR(36) PRIMARY KEY,
    content TEXT NOT NULL
  );

-- Media table
