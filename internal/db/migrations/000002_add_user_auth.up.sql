
CREATE TABLE "user" (
  "id" uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
  "name" varchar(255),
  "dob" date,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp DEFAULT (now())
);

CREATE TABLE "user_auth" (
  "username" varchar PRIMARY KEY,
  "hashed_password" varchar,
  "user_id" uuid,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp DEFAULT (now())
);


CREATE INDEX ON "user_auth" ("user_id");

ALTER TABLE "user_auth" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id");
