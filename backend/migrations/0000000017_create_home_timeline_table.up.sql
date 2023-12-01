BEGIN;

CREATE TABLE SOC_APP_HOME_TIMELINE(
	id bigserial,
	user_id bigint,
	post_id bigint,
	created_at timestamp NOT NULL,
	
	PRIMARY KEY (id),
	FOREIGN KEY(user_id) REFERENCES SOC_APP_USERS(user_id),
	FOREIGN KEY(post_id) REFERENCES SOC_APP_POSTS(post_id)
);
CREATE UNIQUE INDEX UQ_SOC_HOME_TIMELINE ON SOC_APP_USERS(user_id,post_id);
CREATE INDEX IDX_SOC_APP_USERS ON SOC_APP_USERS (created_at);

COMMIT;