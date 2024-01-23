-- Modify "user_roles" table
ALTER TABLE "public"."user_roles" DROP CONSTRAINT "fk_user_roles_role", DROP CONSTRAINT "fk_user_roles_user", ADD
 CONSTRAINT "fk_user_roles_role" FOREIGN KEY ("role_id") REFERENCES "public"."roles" ("id") ON UPDATE CASCADE ON DELETE CASCADE, ADD
 CONSTRAINT "fk_user_roles_user" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("id") ON UPDATE CASCADE ON DELETE CASCADE;
