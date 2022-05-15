CREATE TABLE "class" (
  "id" SERIAL PRIMARY KEY,
  "name" varchar(50) NOT NULL
);

CREATE TABLE "user" (
  "id" SERIAL PRIMARY KEY,
  "username" varchar(50) UNIQUE NOT NULL,
  "password" varchar(60) NOT NULL,
  "name" varchar(75) NOT NULL,
  "surname" varchar(50) NOT NULL,
  "email" varchar(40) NOT NULL,
  "type" varchar(10)
);

CREATE TABLE "admin" (
  "id" int PRIMARY KEY
);

CREATE TABLE "student" (
  "id" int PRIMARY KEY,
  "year" int,
  "status" varchar(10) NOT NULL,
  "class_id" int
);

CREATE TABLE "teacher" (
  "id" int PRIMARY KEY
);

CREATE TABLE "subject" (
  "id" SERIAL PRIMARY KEY,
  "name" varchar(50) NOT NULL
);

CREATE TABLE "teacher_subject_class" (
  "teacher_id" int,
  "subject_id" int,
  "class_id" int
);

CREATE TABLE "activity" (
  "id" SERIAL PRIMARY KEY,
  "title" varchar(100) NOT NULL,
  "description" varchar(1024) NOT NULL
);

CREATE TABLE "activity_subject_class" (
  "activity_id" int,
  "subject_id" int,
  "class_id" int
);

CREATE TABLE "news" (
  "id" SERIAL PRIMARY KEY,
  "title" varchar(50) NOT NULL,
  "description" varchar(1024) NOT NULL
);

CREATE TABLE "news_subject_class" (
  "news_id" int,
  "subject_id" int,
  "class_id" int
);

ALTER TABLE "admin" ADD FOREIGN KEY ("id") REFERENCES "user" ("id");

ALTER TABLE "student" ADD FOREIGN KEY ("id") REFERENCES "user" ("id");

ALTER TABLE "student" ADD FOREIGN KEY ("class_id") REFERENCES "class" ("id");

ALTER TABLE "teacher" ADD FOREIGN KEY ("id") REFERENCES "user" ("id");

ALTER TABLE "teacher_subject_class" ADD FOREIGN KEY ("teacher_id") REFERENCES "teacher" ("id");

ALTER TABLE "teacher_subject_class" ADD FOREIGN KEY ("subject_id") REFERENCES "subject" ("id");

ALTER TABLE "teacher_subject_class" ADD FOREIGN KEY ("class_id") REFERENCES "class" ("id");

ALTER TABLE "activity_subject_class" ADD FOREIGN KEY ("activity_id") REFERENCES "activity" ("id");

ALTER TABLE "activity_subject_class" ADD FOREIGN KEY ("subject_id") REFERENCES "subject" ("id");

ALTER TABLE "activity_subject_class" ADD FOREIGN KEY ("class_id") REFERENCES "class" ("id");

ALTER TABLE "news_subject_class" ADD FOREIGN KEY ("news_id") REFERENCES "news" ("id");

ALTER TABLE "news_subject_class" ADD FOREIGN KEY ("subject_id") REFERENCES "subject" ("id");

ALTER TABLE "news_subject_class" ADD FOREIGN KEY ("class_id") REFERENCES "class" ("id");
