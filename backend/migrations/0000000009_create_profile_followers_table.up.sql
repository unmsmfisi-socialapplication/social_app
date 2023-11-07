BEGIN;

CREATE TABLE SOC_APP_USER_PROFILE_FOLLOW (
  follow_id bigserial,
  follower_profile_id bigint NOT NULL,
  following_profile_id bigint NOT NULL,
  follow_date timestamp,
  PRIMARY KEY (follow_id),
  FOREIGN KEY (follower_profile_id) REFERENCES SOC_APP_USER_PROFILE(profile_id),
  FOREIGN KEY (following_profile_id) REFERENCES SOC_APP_USER_PROFILE(profile_id)
);
CREATE UNIQUE INDEX UQ_SOC_APP_USER_PROFILE_FOLLOW ON SOC_APP_USER_PROFILE_FOLLOW(follower_profile_id,following_profile_id);
CREATE INDEX IDX_FOLLOW_DATE ON SOC_APP_USER_PROFILE_FOLLOW(follow_date);

COMMIT;
