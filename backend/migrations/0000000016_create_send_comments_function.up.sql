BEGIN;

CREATE OR REPLACE  FUNCTION FN_SOC_APP_GET_COMMENTS(p_fec_ini VARCHAR, p_fec_fin VARCHAR)
RETURNS TABLE (
	insertion_date TIMESTAMP,
	comment_id BIGINT,
	post_id BIGINT,
	user_name VARCHAR,
	comment VARCHAR,
	name VARCHAR,
	last_name VARCHAR,
	genre VARCHAR,
	country VARCHAR,
	city VARCHAR
) AS $$
DECLARE
  v_fecha_inicio_proceso TIMESTAMP;
  v_fecha_fin_proceso TIMESTAMP;
BEGIN
  -- Initialize v_fecha_inicio_proceso before running the query
  v_fecha_inicio_proceso := now();

  RETURN QUERY 
  SELECT
	a.insertion_date comment_date,
	a.comment_id,
	a.post_id,
	a.comment,
	b.user_name,
	c.name,
	c.last_name,
  	c.genre,
  	c.country,
  	c.city
  FROM soc_app_posts_comments a
  INNER JOIN soc_app_users b
    ON a.user_id = b.user_id
  INNER JOIN soc_app_user_profile c
  	ON a.user_id = c.user_id
  WHERE
    a.insertion_date >= TO_TIMESTAMP(p_fec_ini, 'YYYY-MM-DD HH24:MI')
    AND a.insertion_date < TO_TIMESTAMP(p_fec_fin, 'YYYY-MM-DD HH24:MI');

  -- Initialize v_fecha_fin_proceso after running the query
  v_fecha_fin_proceso := now();
  INSERT INTO soc_app_seg_get_comments_data(
	start_proc_date,
	end_proc_date,
	query_fec_ini,
	query_fec_fin
  )
  VALUES(
	v_fecha_inicio_proceso,
	v_fecha_fin_proceso,
	TO_TIMESTAMP(p_fec_ini, 'YYYY-MM-DD HH24:MI'),
	TO_TIMESTAMP(p_fec_fin, 'YYYY-MM-DD HH24:MI')
  );
END;
$$ LANGUAGE plpgsql;

COMMIT;
