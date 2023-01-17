CREATE DATABASE IF NOT EXISTS devbook;

USE devbook;

CREATE TABLE IF NOT EXISTS users(
  id int auto_increment primary key,
  name varchar(255) not null,
  nick varchar(32) not null unique,
  email varchar(320) not null unique,
  password varchar(320) not null,
  createdAt timestamp not null default current_timestamp()
) ENGINE=INNODB;