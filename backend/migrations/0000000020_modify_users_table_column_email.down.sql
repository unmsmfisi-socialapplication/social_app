BEGIN;

ALTER TABLE soc_app_users
ALTER COLUMN email DROP NOT NULL;

COMMIT;