BEGIN;

ALTER TABLE soc_app_users
ALTER COLUMN insertion_date SET NOT NULL;

COMMIT;