CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE "author" (
  "id" uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
  "name" varchar NOT NULL,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp DEFAULT (now()),
  CONSTRAINT "unique_author_id" UNIQUE( "id" )
);

CREATE TABLE "publisher" (
  "id" uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
  "name" varchar NOT NULL,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp DEFAULT (now()),
  CONSTRAINT "unique_publisher_id" UNIQUE( "id" )
);

CREATE TABLE "book" (
  "id" uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
  "title" varchar NOT NULL,
  "description" varchar,
  "year" varchar,
  "author_id" uuid,
  "publisher_id" uuid,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp DEFAULT (now()),
  CONSTRAINT "unique_book_id" UNIQUE( "id" )
);

CREATE INDEX ON "author" ("name");
CREATE INDEX ON "book" ("author_id");
CREATE INDEX ON "book" ("publisher_id");


ALTER TABLE "book" ADD FOREIGN KEY ("author_id") REFERENCES "author" ("id");

ALTER TABLE "book" ADD FOREIGN KEY ("publisher_id") REFERENCES "publisher" ("id");