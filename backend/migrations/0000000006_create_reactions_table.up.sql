BEGIN;

CREATE TABLE SOC_APP_POSTS_REACTIONS (
  post_reaction_id bigserial,
  reaction_id bigserial,
  user_id bigint NOT NULL,
  post_id bigint NOT NULL,
  insertion_date timestamp NOT NULL,
  update_date timestamp NOT NULL,
  PRIMARY KEY (post_reaction_id),
  FOREIGN KEY (user_id) REFERENCES SOC_APP_USERS(user_id),
  FOREIGN KEY (post_id) REFERENCES SOC_APP_POSTS(post_id),
  FOREIGN KEY (reaction_id) REFERENCES SOC_APP_M_USERS_REACTIONS(reaction_id)
);
CREATE UNIQUE INDEX UQ_SOC_APP_POSTS_REACTIONS ON SOC_APP_POSTS_REACTIONS(user_id,post_id);

COMMIT;