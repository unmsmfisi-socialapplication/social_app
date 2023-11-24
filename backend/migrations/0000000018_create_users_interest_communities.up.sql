BEGIN;

CREATE TABLE SOC_APP_USERS_INTEREST_COMMUNITIES
(
  user_interest_community_id bigserial NOT NULL,
  user_id bigint NOT NULL,
  community_id bigint NOT NULL,
  PRIMARY KEY(user_interest_community_id),
  FOREIGN KEY(user_id) REFERENCES SOC_APP_USERS(user_id),
  FOREIGN KEY(community_id) REFERENCES SOC_APP_INTEREST_COMMUNITIES(community_id)
);

CREATE UNIQUE INDEX UQ_SOC_APP_USERS_INTEREST_COMMUNITIES ON SOC_APP_USERS_INTEREST_COMMUNITIES(user_id, community_id);

COMMIT;