BEGIN;

CREATE TABLE SOC_APP_POSTS(
  post_id bigserial,
  user_id bigint,
  title VARCHAR(100) NOT NULL,
  description VARCHAR(1000),
  has_multimedia boolean,
  public boolean,
  multimedia text,
  insertion_date timestamp NOT NULL,
  update_date timestamp NOT NULL,
  PRIMARY KEY (post_id),
  FOREIGN KEY (user_id) REFERENCES SOC_APP_USERS(user_id)
);
CREATE INDEX IDX1_SOC_APP_POSTS ON SOC_APP_POSTS(insertion_date);
CREATE INDEX IDX2_SOC_APP_POSTS ON SOC_APP_POSTS(update_date);

COMMIT;