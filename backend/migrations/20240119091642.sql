-- Create "users" table
CREATE TABLE "public"."users" (
  "id" bigserial NOT NULL,
  "username" text NULL,
  "password" text NULL,
  "email" text NULL,
  PRIMARY KEY ("id")
);
-- Create "books" table
CREATE TABLE "public"."books" (
  "id" bigserial NOT NULL,
  "name" text NULL,
  "author" text NULL,
  "description" text NULL,
  "photo" text NULL,
  "user_id" bigint NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_users_books" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create "permissions" table
CREATE TABLE "public"."permissions" (
  "id" bigserial NOT NULL,
  "permission" text NULL,
  PRIMARY KEY ("id")
);
-- Create "roles" table
CREATE TABLE "public"."roles" (
  "id" bigserial NOT NULL,
  "name" text NULL,
  PRIMARY KEY ("id")
);
-- Create "role_permissions" table
CREATE TABLE "public"."role_permissions" (
  "id" bigserial NOT NULL,
  "role_id" bigint NULL,
  "permission_id" bigint NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_role_permissions_permission" FOREIGN KEY ("permission_id") REFERENCES "public"."permissions" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "fk_role_permissions_role" FOREIGN KEY ("role_id") REFERENCES "public"."roles" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create "user_roles" table
CREATE TABLE "public"."user_roles" (
  "role_id" bigint NOT NULL,
  "user_id" bigint NOT NULL,
  PRIMARY KEY ("role_id", "user_id"),
  CONSTRAINT "fk_user_roles_role" FOREIGN KEY ("role_id") REFERENCES "public"."roles" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "fk_user_roles_user" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
