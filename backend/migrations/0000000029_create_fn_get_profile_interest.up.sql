BEGIN;


CREATE OR REPLACE FUNCTION FN_SOC_APP_GET_PROFILE_INTEREST(
    p_profile_id BIGINT,
    p_page_size INT,
    p_page_number INT
)
RETURNS TABLE (
    interest_id BIGINT,
    user_id BIGINT,
    profile_id BIGINT,
    interest_name VARCHAR,
    user_name VARCHAR
) AS $$
DECLARE
BEGIN

    RETURN QUERY 
    SELECT
        a.interest_id,
        c.user_id,
        d.profile_id,
        b.interest_name,
        c.user_name
    FROM soc_app_users_interest_topics a
    INNER JOIN soc_app_m_users_interests b
        ON b.interest_id = a.interest_id
    INNER JOIN soc_app_users c
        ON c.user_id = a.user_id
    INNER JOIN soc_app_user_profile d
        ON a.user_id = d.user_id
    WHERE d.profile_id = p_profile_id
    LIMIT p_page_size
    OFFSET (p_page_number - 1) * p_page_size;
END;
$$ LANGUAGE plpgsql;


COMMIT;
