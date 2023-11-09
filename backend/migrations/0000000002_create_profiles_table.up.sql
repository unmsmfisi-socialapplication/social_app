BEGIN;

CREATE TABLE SOC_APP_USER_PROFILE(
  profile_id bigserial,
  user_id bigint NOT NULL,
  birth_date date NOT NULL,
  name VARCHAR(500) NOT NULL,
  last_name VARCHAR(500) NOT NULL,
  about_me text,
  genre varchar(10) NOT NULL,
  address VARCHAR(500),
  country VARCHAR(50) NOT NULL,
  city VARCHAR(50) NOT NULL,
  insertion_date timestamp,
  update_date timestamp,
  profile_picture text,
  PRIMARY KEY (profile_id),
  FOREIGN KEY (user_id) REFERENCES SOC_APP_USERS(user_id)
);
CREATE UNIQUE INDEX UQ_SOC_APP_USER_PROFILE ON SOC_APP_USER_PROFILE(user_id);
CREATE INDEX IDX1_SOC_APP_USER_PROFILE ON SOC_APP_USER_PROFILE(insertion_date);
CREATE INDEX IDX2_SOC_APP_USER_PROFILE ON SOC_APP_USER_PROFILE(update_date);

COMMIT;