CREATE TABLE SOC_APP_POSTS_COMMENTS (
  comment_id bigserial,
  user_id bigint NOT NULL,
  post_id bigint NOT NULL,
  comment varchar(1000) NOT NULL,
  insertion_date timestamp NOT NULL,
  update_date timestamp NOT NULL,
  parent_comment_id bigint,
  PRIMARY KEY (comment_id),
  FOREIGN KEY (user_id) REFERENCES SOC_APP_USERS(user_id),
  FOREIGN KEY (post_id) REFERENCES SOC_APP_POSTS(post_id),
  FOREIGN KEY (parent_comment_id) REFERENCES SOC_APP_POSTS_COMMENTS(comment_id)
);


CREATE INDEX IDX1_SOC_APP_POSTS_COMMENTS ON SOC_APP_POSTS_COMMENTS(insertion_date);
CREATE INDEX IDX2_SOC_APP_POSTS_COMMENTS ON SOC_APP_POSTS_COMMENTS(update_date);

COMMIT;