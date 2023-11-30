BEGIN;

CREATE TABLE soc_app_seg_get_profiles_by_name
(
    search_profile_id bigserial,
    start_proc_date timestamp,
    end_proc_date timestamp,
    keyword VARCHAR(500),

    CONSTRAINT PK_SOC_APP_SEG_GET_PROFILES_BY_NAME PRIMARY KEY (search_profile_id)       

);

CREATE INDEX idx_soc_app_seg_get_profiles_by_name ON soc_app_seg_get_profiles_by_name (start_proc_date, keyword);

COMMIT;