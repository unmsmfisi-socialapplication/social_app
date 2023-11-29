-- PROCEDURE: bi.process_bi_spam()

--DROP PROCEDURE IF EXISTS bi.process_bi_spam();

CREATE OR REPLACE PROCEDURE bi.process_bi_spam(
	)
LANGUAGE 'plpgsql'
AS $BODY$
BEGIN
    -- Paso 1: Vaciar la tabla bi_sentimental_transformada
    TRUNCATE TABLE bi.bi_spam_transformada;
    
    -- Paso 2: Transformar y limpiar los datos
    -- Esto incluirá eliminar duplicados, nulos, blancos y caracteres no deseados
    -- También se asegurará de que la columna 'category' tenga solo los valores -1, 0, 1
    INSERT INTO bi.bi_spam_transformada ("Post","Prediction","fecha_sp")
    SELECT DISTINCT
        TRIM(REGEXP_REPLACE("Post", '[^a-zA-Z0-9ÁÉÍÓÚáéíóúñÑ\s]', '', 'g')),
        CASE WHEN "Prediction" IN ('-1', '0', '1') THEN "Prediction" ELSE NULL END,
        CURRENT_TIMESTAMP
    FROM bi.bi_spam
    WHERE "Post" IS NOT NULL AND TRIM("Post") != '';

    -- Eliminar las filas con category NULL después de la limpieza
    DELETE FROM bi.bi_spam_transformada WHERE "Post" IS NULL;
END;
$BODY$;
ALTER PROCEDURE bi.process_bi_spam()
    OWNER TO useradmin;

