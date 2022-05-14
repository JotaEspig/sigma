/*
link for diagram: https://dbdiagram.io/d/62616cbe1072ae0b6ac6170f

Contains the query for the creation of the tables
To configure the database correctly, you must create these environment variables:
- DB_USERNAME
- DB_PASSWORD
- DB_HOST
- DB_PORT

And you must run these commands:
1- Create a database and add its name to a env variable called 'DB_DB'
2- Access the database (\c <dbname>)
3- Run these commands (you can just copy and paste)

*If you're using Docker, you don't need to create the env variable, because it
already exists. So you just need to access the shell of the database container
and do the rest

** The database is not complete ** 

TODO jota: it's needed to separate the table user to "admin", "teacher", "student"

*/

CREATE TABLE "user" (
  "id" SERIAL PRIMARY KEY,
  "username" varchar(50) UNIQUE NOT NULL,
  "password" varchar(60) NOT NULL,
  "email" varchar(40) NOT NULL,
  "name" varchar(100) NOT NULL
);

CREATE TABLE "user_class" (
  "user_id" int,
  "class_id" int
);

CREATE TABLE "class" (
  "id" SERIAL PRIMARY KEY,
  "subject" varchar(50) NOT NULL
);

ALTER TABLE "user_class" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id");

ALTER TABLE "user_class" ADD FOREIGN KEY ("class_id") REFERENCES "class" ("id");
