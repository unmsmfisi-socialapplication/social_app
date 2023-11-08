BEGIN;

CREATE OR REPLACE FUNCTION fn_soc_app_get_posts(p_fec_ini VARCHAR, p_fec_fin VARCHAR)
RETURNS TABLE (
  user_name VARCHAR,
  name VARCHAR,
  last_name VARCHAR,
  genre VARCHAR,
  country VARCHAR,
  city VARCHAR,
  post_id BIGINT,
  title VARCHAR,
  description VARCHAR,
  post_date TIMESTAMP
) AS $$
DECLARE
  v_fecha_inicio_proceso TIMESTAMP;
  v_fecha_fin_proceso TIMESTAMP;
BEGIN
  -- Initialize v_fecha_inicio_proceso before running the query
  v_fecha_inicio_proceso := now();

  RETURN QUERY 
  SELECT
    b.user_name,
    c.name,
    c.last_name,
    c.genre,
    c.country,
    c.city,
    a.post_id,
    a.title,
    a.description,
    a.insertion_date post_date
  FROM soc_app_posts a
  INNER JOIN soc_app_users b
    ON a.user_id = b.user_id
  INNER JOIN soc_app_user_profile c
    ON c.user_id = a.user_id
  WHERE
    a.insertion_date >= TO_TIMESTAMP(p_fec_ini, 'YYYY-MM-DD HH24:MI')
    AND a.insertion_date < TO_TIMESTAMP(p_fec_fin, 'YYYY-MM-DD HH24:MI');

  -- Initialize v_fecha_fin_proceso after running the query
  v_fecha_fin_proceso := now();
  INSERT INTO soc_app_seg_get_posts_data(
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