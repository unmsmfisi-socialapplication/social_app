/*
@Description: 
	This table performs the test that simulates changes to BI process tables that only 
	show a satisfactory message if it is INSERT or DELETE, 
	if it is UPDATE an Error message will appear. 
	This is because it is good practice for BI processes.
@author: Gregory Sánchez
*/

/*
Creation Table Test_Log
*/
CREATE TABLE IF NOT EXISTS bi.test_tg
(
    campo1_id serial,
    campo2 text,
    campo3 bigint,
    CONSTRAINT log_pkey PRIMARY KEY (campo1_id)
);

/*
Creation Table Log
*/
CREATE TABLE bi.bi_log (
    log_id serial PRIMARY KEY,
    event_type text,
    event_description text,
	event_tablename text,
    event_timestamp timestamp,
    username name,
	status text,
    json_data text
);

/*
Creation Function Trigger
*/
CREATE OR REPLACE FUNCTION bi.fntg_bi_log()
RETURNS TRIGGER AS $$
DECLARE
	username name := current_user;
	mensaje_error text := '';
BEGIN
	
	BEGIN
		IF TG_OP IN ('DELETE') THEN
			INSERT INTO bi.bi_log (event_type, event_description, event_tablename, event_timestamp, username, status, json_data)
			VALUES (TG_OP, 'Successfull', TG_TABLE_NAME, now(), username, '', OLD::text);
		ELSIF TG_OP IN ('INSERT') THEN
			INSERT INTO bi.bi_log (event_type, event_description, event_tablename, event_timestamp, username, status, json_data)
			VALUES (TG_OP, 'Successfull', TG_TABLE_NAME, now(), username, '', NEW::text);
		ELSIF TG_OP IN ('UPDATE') THEN
			INSERT INTO bi.bi_log (event_type, event_description, event_tablename, event_timestamp, username, status, json_data)
			VALUES (TG_OP, 'Unsupported operation', TG_TABLE_NAME, now(), username, 'E', OLD::text ||', '||NEW::text);

			RETURN NULL;
		END IF;
	EXCEPTION
        WHEN OTHERS THEN
            mensaje_error := SQLERRM;
            INSERT INTO bi.bi_log (event_type, event_description, event_tablename, event_timestamp, username, status, json_data)
            VALUES (TG_OP, mensaje_error, TG_TABLE_NAME, now(), username, 'E', NEW::text);
    END;
	
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

/*
Creation Trigger of the table
*/
CREATE TRIGGER tg_after_test_tg
AFTER INSERT OR UPDATE OR DELETE ON bi.test_tg
FOR EACH ROW
EXECUTE FUNCTION bi.fntg_bi_log();

/*
Test
*/
-- Insert to table test_log
INSERT INTO bi.test_tg (campo2, campo3)
VALUES ('Proceso ETL 1', 241);
INSERT INTO bi.test_tg (campo2, campo3)
VALUES ('Proceso ETL 2', 542);
INSERT INTO bi.test_tg (campo2, campo3)
VALUES ('Proceso ETL 3', 843);

-- Message -> Successfull
DELETE FROM bi.test_tg;

-- Message -> Unsupported operation
UPDATE bi.test_tg
SET campo2='CAMBIADO'
where campo1_id=7;

--Test
select * from bi.bi_log order by log_id desc;