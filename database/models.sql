\c luis;
DROP DATABASE students;

CREATE DATABASE students;
\c students;

CREATE TABLE students (
  id SERIAL PRIMARY KEY,
  name TEXT NOT NULL,
  surname TEXT NOT NULL,
  code TEXT DEFAULT '',
  grade TEXT DEFAULT '',
  birthdate DATE DEFAULT '1900-01-01',
  public_id TEXT DEFAULT '',
  photo TEXT DEFAULT ''
);

INSERT INTO students (name, surname) 
  VALUES ('Student Name', 'Student Lastname');

