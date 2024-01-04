-- Modify "books" table
ALTER TABLE "public"."books" DROP COLUMN "created_at", DROP COLUMN "updated_at", DROP COLUMN "deleted_at";
-- Modify "users" table
ALTER TABLE "public"."users" DROP COLUMN "created_at", DROP COLUMN "updated_at", DROP COLUMN "deleted_at";
