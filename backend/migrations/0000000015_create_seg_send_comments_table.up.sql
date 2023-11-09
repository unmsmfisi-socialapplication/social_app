CREATE TABLE SOC_APP_SEG_GET_COMMENTS_DATA(
	id_query bigserial,
	start_proc_date timestamp,
	end_proc_date timestamp,
	query_fec_ini timestamp,
	query_fec_fin timestamp
);
CREATE UNIQUE INDEX UQ_SOC_APP_SEG_GET_COMMENTS_DATA ON SOC_APP_SEG_GET_COMMENTS_DATA(id_query);
CREATE INDEX IDX1_SOC_APP_SEG_GET_COMMENTS_DATA ON SOC_APP_SEG_GET_COMMENTS_DATA (start_proc_date);

COMMIT;