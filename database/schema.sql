\c luis;
DROP DATABASE students;

CREATE DATABASE students;
\c students;

CREATE TYPE grade_type AS ENUM (
  '',
  'sala de tres',
  'sala de cuatro', 
  'sala de cinco', 

  'primer grado', 
  'segundo grado',
  'tercer grado',
  'cuarto grado',
  'quinto grado',
  'sexto grado',

  'primer año',
  'segundo año',
  'tercer año',
  'cuarto año',
  'quinto año'
  );

CREATE TYPE section_type AS ENUM ('','A','B','C','D','E','F','G','H','I','J','K','L','M','N','O','P','Q','R','S','T','U','V','W','X','Y','Z');

CREATE TABLE students (
  id SERIAL PRIMARY KEY,
  name TEXT NOT NULL,
  surname TEXT NOT NULL,
  code TEXT DEFAULT '',
  grade grade_type DEFAULT '',
  section section_type DEFAULT '',
  birthdate DATE DEFAULT '0001-01-01',
  public_id TEXT DEFAULT '',
  photo TEXT DEFAULT ''
);