CREATE DATABASE IF NOT EXISTS devbook;

USE devbook;

DROP TABLE IF EXISTS posts;

DROP TABLE IF EXISTS followers;

DROP TABLE IF EXISTS users;

CREATE TABLE IF NOT EXISTS users(
  id int auto_increment primary key,
  name varchar(255) not null,
  nick varchar(32) not null unique,
  email varchar(320) not null unique,
  password varchar(320) not null,
  createdAt timestamp not null default current_timestamp()
) ENGINE = INNODB;

CREATE TABLE IF NOT EXISTS followers(
  user_id int not null,
  FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
  follower_id int not null,
  FOREIGN KEY (follower_id) REFERENCES users(id) ON DELETE CASCADE,
  primary key (user_id, follower_id)
) ENGINE = INNODB;

CREATE TABLE IF NOT EXISTS posts(
  id int auto_increment primary key,
  title varchar(255) not null,
  content varchar(4096) not null,
  author_id int not null,
  FOREIGN KEY (author_id) REFERENCES users(id) ON DELETE CASCADE,
  likes int default 0,
  createdAt timestamp not null default current_timestamp()
) ENGINE = INNODB;