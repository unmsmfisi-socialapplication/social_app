BEGIN;

CREATE TABLE sa.SOC_APP_USER_PROFILE_FOLLOW (
  follow_id bigserial,
  follower_user_id bigint NOT NULL,
  following_user_id bigint NOT NULL,
  follow_date timestamp,
  PRIMARY KEY (follow_id),
  FOREIGN KEY (follower_user_id) REFERENCES sa.SOC_APP_USER_PROFILE(profile_id),
  FOREIGN KEY (following_user_id) REFERENCES sa.SOC_APP_USER_PROFILE(profile_id)
);
CREATE UNIQUE INDEX UQ_FOLLOW_DATE ON sa.SOC_APP_USER_PROFILE_FOLLOW(follower_user_id,following_user_id);
CREATE INDEX IDX_FOLLOW_DATE ON sa.SOC_APP_USER_PROFILE_FOLLOW(follow_date);

COMMIT;