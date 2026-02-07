-- Auth Table
CREATE TABLE auth_sessions(
  id VARCHAR(36) PRIMARY KEY,
  user_id VARCHAR(36) NOT NULL UNIQUE,
  email VARCHAR(255) NOT NULL UNIQUE,
  password_hash TEXT NOT NULL,
  is_active INTEGER DEFAULT 1,
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);
-- Users table
CREATE TABLE users (
  id VARCHAR(36) PRIMARY KEY,
  username VARCHAR(50) UNIQUE,
  display_name VARCHAR(255),
  account_status VARCHAR(20) NOT NULL DEFAULT 'active',
  username_changed_at DATETIME NULL,
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  deleted_at DATETIME NULL,
  deleted_after INTEGER DEFAULT 30
);
-- Guilds table
CREATE TABLE guilds (
  id VARCHAR(36) PRIMARY KEY,
  guild_name VARCHAR(255) NOT NULL,
  invite_link VARCHAR(255)
);
-- User Guilds joined table
CREATE TABLE user_guilds (
  user_id VARCHAR(36) NOT NULL,
  guild_id VARCHAR(36) NOT NULL,
  role VARCHAR(50) DEFAULT 'member',
  joined_at DATETIME DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (user_id, guild_id),
  FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
  FOREIGN KEY (guild_id) REFERENCES guilds(id) ON DELETE CASCADE
);
-- Messages table
CREATE TABLE messages (
  id VARCHAR(36) PRIMARY KEY,
  content TEXT NOT NULL
);
-- Media table