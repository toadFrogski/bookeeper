-- Create "checkouts" table
CREATE TABLE "public"."checkouts" (
  "id" bigserial NOT NULL,
  "book_id" bigint NULL,
  "user_id" bigint NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_checkouts_book" FOREIGN KEY ("book_id") REFERENCES "public"."books" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "fk_checkouts_user" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
