BEGIN;

CREATE OR REPLACE FUNCTION FN_SOC_APP_GET_PROFILES_BY_NAME(
    p_keyword VARCHAR,
    p_page_size INT,
    p_page_number INT
)
RETURNS TABLE (
    profile_id bigint,
    user_id bigint,
    name VARCHAR,
    last_name VARCHAR,
    about_me text,
    genre VARCHAR,
    address VARCHAR,
    country VARCHAR,
    city VARCHAR,
    insertion_date TIMESTAMP,
    update_date TIMESTAMP,
    profile_picture TEXT
) AS $$
DECLARE
    v_fecha_inicio_proceso TIMESTAMP;
    v_fecha_fin_proceso TIMESTAMP;
BEGIN
    v_fecha_inicio_proceso := now();

    RETURN QUERY 
        SELECT
            a.profile_id,
            a.user_id,
            a.name,
            a.last_name,
            a.about_me,
            a.genre,
            a.address,
            a.country,
            a.city,
            a.insertion_date,
            a.update_date,
            a.profile_picture
        FROM soc_app_user_profile a
        WHERE UPPER(a.name) LIKE UPPER('%'||p_keyword||'%')
        ORDER BY a.insertion_date DESC
        LIMIT p_page_size
        OFFSET (p_page_number - 1) * p_page_size;

    v_fecha_fin_proceso := now();
    INSERT INTO soc_app_seg_get_PROFILES_BY_NAME(
        start_proc_date,
        end_proc_date,
        keyword
    )
    VALUES(
        v_fecha_inicio_proceso,
        v_fecha_fin_proceso,
        p_keyword
    );
END;
$$ LANGUAGE plpgsql;


COMMIT;