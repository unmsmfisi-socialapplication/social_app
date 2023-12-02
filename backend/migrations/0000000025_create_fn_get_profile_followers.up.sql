CREATE OR REPLACE FUNCTION FN_SOC_APP_GET_PROFILE_FOLLOWERS(
    p_profile_id BIGINT,
    p_page_size INT,
    p_page_number INT
)
RETURNS TABLE (
	profile_id BIGINT,
	user_id BIGINT,
	name VARCHAR,
	last_name VARCHAR,
	profile_picture TEXT,
	user_name VARCHAR
) AS $$
DECLARE
BEGIN

    RETURN QUERY 
    SELECT
        b.profile_id,
        c.user_id,
        b.name,
        b.last_name,
        b.profile_picture,
        c.user_name
    FROM soc_app_user_profile_follow a 
    INNER JOIN soc_app_user_profile b
        ON a.following_profile_id = b.profile_id
    INNER JOIN soc_app_users c
        ON b.user_id = c.user_id
    WHERE
        following_profile_id = p_profile_id
    LIMIT p_page_size
    OFFSET (p_page_number - 1) * p_page_size;
END;
$$ LANGUAGE plpgsql;