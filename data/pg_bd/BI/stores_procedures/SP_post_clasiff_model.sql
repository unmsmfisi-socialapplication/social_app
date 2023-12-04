CREATE OR REPLACE PROCEDURE bi.get_data_postclasiff_model(
	IN name_schema text,
	IN name_table text,
	IN model_ver text)
LANGUAGE 'plpgsql'
AS $BODY$
DECLARE 
"ID" INT;
feh_create TIMESTAMP;
name_user TEXT;
BEGIN
SELECT TO_CHAR(CURRENT_DATE, 'YYYYMMDD')::INT into "ID";

SELECT CURRENT_TIMESTAMP INTO feh_create;

SELECT current_user INTO name_user;
EXECUTE 
'INSERT INTO bi.resultspostclasif (ID, post_description, category,model_version,feh_create,name_user) 
SELECT ID ,post_text,label,feh_create ,name_user FROM' ||name_schema|| '.' ||name_table
USING 
"ID",
feh_create,
name_user;
END;
$BODY$;
