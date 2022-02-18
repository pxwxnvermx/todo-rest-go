CREATE TABLE "todos" (
  "id" bigserial PRIMARY KEY,
  "description" varchar NOT NULL,
  "is_completed" boolean NOT NULL DEFAULT false,
  "owner" bigint,
  "created_at" timestamp,
  "updated_at" timestamp,
  "deleted_at" timestamp
);

CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "email" varchar UNIQUE NOT NULL,
  "password" varchar,
  "created_at" timestamp,
  "updated_at" timestamp,
  "deleted_at" timestamp
);

ALTER TABLE "todos" ADD FOREIGN KEY ("owner") REFERENCES "users" ("id");

CREATE INDEX ON "todos" ("owner");

CREATE INDEX ON "users" ("email");
