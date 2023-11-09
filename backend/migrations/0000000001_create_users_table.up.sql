BEGIN;

CREATE TABLE SOC_APP_USERS(
  user_id bigserial,
  insertion_date timestamp,
  phone bigint,
  email  varchar(255),
  user_name varchar(25),
  password text,
  CONSTRAINT PK_SOC_APP_USERS PRIMARY KEY (user_id)
);
CREATE UNIQUE INDEX UQ_SOC_APP_USERS ON SOC_APP_USERS(phone,email,user_name);
CREATE INDEX IDX_SOC_APP_USERS ON SOC_APP_USERS (insertion_date);

COMMIT;