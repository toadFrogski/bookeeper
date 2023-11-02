-- Create "users" table
CREATE TABLE "public"."users" (
  "username" text NULL,
  "password" text NULL,
  "email" text NULL,
  "id" bigserial NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  PRIMARY KEY ("id")
);
