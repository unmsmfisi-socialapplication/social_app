BEGIN;

ALTER TABLE SOC_APP_USERS
ADD COLUMN name varchar(255);

ALTER TABLE SOC_APP_USERS
DROP COLUMN phone;

ALTER TABLE SOC_APP_USERS
ADD COLUMN phone varchar(255);

ALTER TABLE SOC_APP_USERS
ADD COLUMN photo TEXT;

COMMIT;