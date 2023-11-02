-- Create "books" table
CREATE TABLE "public"."books" (
  "id" bigserial NOT NULL,
  "name" text NULL,
  "author" text NULL,
  "description" text NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  PRIMARY KEY ("id")
);
