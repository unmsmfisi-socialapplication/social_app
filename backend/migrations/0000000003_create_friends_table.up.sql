BEGIN;

CREATE TABLE sa.SOC_APP_FRIENDS (
  friend_id bigserial,
  user_id1 bigint,
  user_id2 bigint,
  friendship_date timestamp,
  PRIMARY KEY (friend_id),
  FOREIGN KEY (user_id1) REFERENCES sa.SOC_APP_USERS(user_id),
  FOREIGN KEY (user_id2) REFERENCES sa.SOC_APP_USERS(user_id)
);
CREATE UNIQUE INDEX UQ_SOC_APP_FRIENDS ON sa.SOC_APP_FRIENDS(user_id1,user_id2);

COMMIT;