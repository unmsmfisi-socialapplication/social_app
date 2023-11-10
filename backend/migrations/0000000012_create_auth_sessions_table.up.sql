CREATE TABLE soc_app_auth_sessions (
    id integer NOT NULL,
    user_name character varying(255) NOT NULL,
    logged boolean,
    "timestamp" timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    jti character varying(255)
);
