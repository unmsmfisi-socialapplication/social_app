BEGIN;

CREATE TABLE SOC_APP_COMMUNITIES(
  community_id bigserial NOT NULL,
  community_name VARCHAR(200),
  community_description VARCHAR(500) DEFAULT '',
  interest_id bigint NOT NULL,
  insertion_date timestamp,
  update_date timestamp,
  PRIMARY KEY (community_id),
  FOREIGN KEY(interest_id) REFERENCES SOC_APP_M_USERS_INTERESTS(interest_id)
);
CREATE UNIQUE INDEX UQ_SOC_APP_COMMUNITIES ON SOC_APP_COMMUNITIES(community_name);
CREATE INDEX IDX1_SOC_APP_COMMUNITIES ON SOC_APP_COMMUNITIES(insertion_date);
CREATE INDEX IDX2_SOC_APP_COMMUNITIES ON SOC_APP_COMMUNITIES(update_date);

COMMIT;