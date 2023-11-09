BEGIN;

CREATE TABLE soc_app_seg_get_posts_data(
  id_query bigserial,
  start_proc_date timestamp,
  end_proc_date timestamp,
  query_fec_ini timestamp,
  query_fec_fin timestamp
);
CREATE UNIQUE INDEX UQ_soc_app_seg_get_posts_data ON soc_app_seg_get_posts_data(id_query);
CREATE INDEX IDX1_soc_app_seg_get_posts_data ON soc_app_seg_get_posts_data (start_proc_date);

COMMIT;
