CREATE TABLE "author" (
  "id" INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "name" varchar NOT NULL,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp DEFAULT (now())
);

CREATE TABLE "publisher" (
  "id" INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "name" varchar NOT NULL,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp DEFAULT (now())
);

CREATE TABLE "book" (
  "id" INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "title" varchar NOT NULL,
  "description" varchar,
  "year" varchar,
  "author_id" int,
  "publisher_id" int
);

CREATE INDEX ON "author" ("name");

CREATE INDEX ON "book" ("author_id");

CREATE INDEX ON "book" ("publisher_id");

ALTER TABLE "book" ADD FOREIGN KEY ("author_id") REFERENCES "author" ("id");

ALTER TABLE "book" ADD FOREIGN KEY ("publisher_id") REFERENCES "publisher" ("id");