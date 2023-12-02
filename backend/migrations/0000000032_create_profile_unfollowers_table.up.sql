BEGIN;

CREATE TABLE SOC_APP_USER_PROFILE_UNFOLLOW (
    unfollow_id bigserial,
    follower_profile_id bigint NOT NULL,
    unfollowed_profile_id bigint NOT NULL,
    unfollow_date timestamp,
    PRIMARY KEY (unfollow_id),
    FOREIGN KEY (follower_profile_id) REFERENCES SOC_APP_USER_PROFILE(profile_id),
    FOREIGN KEY (unfollowed_profile_id) REFERENCES SOC_APP_USER_PROFILE(profile_id)
);

CREATE UNIQUE INDEX UQ_SOC_APP_USER_PROFILE_UNFOLLOW ON SOC_APP_USER_PROFILE_UNFOLLOW(follower_profile_id, unfollowed_profile_id);

CREATE INDEX IDX_SOC_APP_USER_PROFILE_UNFOLLOW ON SOC_APP_USER_PROFILE_UNFOLLOW(unfollow_date);

COMMIT;