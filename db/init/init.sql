CREATE DATABASE public;

CREATE TABLE Users (
	ID SERIAL primary key,
	username text,
	password text
);