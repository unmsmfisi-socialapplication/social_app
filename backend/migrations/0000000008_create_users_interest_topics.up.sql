BEGIN;

CREATE TABLE sa.SOC_APP_USERS_INTEREST_TOPICS
(
  user_interest_id bigserial,
  user_id bigint NOT NULL,
  interest_id bigint NOT NULL,
  PRIMARY KEY(user_interest_id),
  FOREIGN KEY(user_id) REFERENCES SOC_APP_USERS(user_id),
  FOREIGN KEY(interest_id) REFERENCES SOC_APP_M_USERS_INTERESTS(interest_id)
);

CREATE UNIQUE INDEX IDX1_SOC_APP_USERS_INTEREST_TOPICS ON SOC_APP_USERS_INTEREST_TOPICS(user_interest_id);

CREATE UNIQUE INDEX IDX2_SOC_APP_USERS_INTEREST_TOPICS ON SOC_APP_USERS_INTEREST_TOPICS(user_id, interest_id);

COMMIT;