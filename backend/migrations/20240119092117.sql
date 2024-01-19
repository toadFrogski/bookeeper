-- Modify "role_permissions" table
ALTER TABLE "public"."role_permissions" DROP CONSTRAINT "role_permissions_pkey", DROP COLUMN "id", ALTER COLUMN "role_id" SET NOT NULL, ALTER COLUMN "permission_id" SET NOT NULL, ADD PRIMARY KEY ("permission_id", "role_id");
