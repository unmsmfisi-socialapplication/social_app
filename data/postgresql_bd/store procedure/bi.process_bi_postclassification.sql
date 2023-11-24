-- PROCEDURE: bi.process_bi_sentimental()

-- DROP PROCEDURE IF EXISTS bi.process_bi_sentimental();

CREATE OR REPLACE PROCEDURE bi.process_bi_postclassification(
	)
LANGUAGE 'plpgsql'
AS $BODY$
BEGIN
    -- Paso 1: Vaciar la tabla bi_sentimental_transformada
    TRUNCATE TABLE bi.bi_postclassification_transformada;
    
    -- Paso 2: Transformar y limpiar los datos
    -- Esto incluirá eliminar duplicados, nulos, blancos y caracteres no deseados
    -- También se asegurará de que la columna 'category' tenga solo los valores -1, 0, 1
    INSERT INTO bi.bi_postclassification_transformada (clean_text, category, fecha_actualizacion)
     SELECT DISTINCT
        TRIM(REGEXP_REPLACE(post_description, '[^a-zA-Z0-9ÁÉÍÓÚáéíóúñÑ\s]', '', 'g')),category,
        CURRENT_TIMESTAMP
    FROM bi.bi_postclassification
    WHERE post_description IS NOT NULL AND TRIM(post_description) != '';

    -- Eliminar las filas con category NULL después de la limpieza
    DELETE FROM bi.bi_postclassification_transformada WHERE category IS NULL;
END;
$BODY$;
ALTER PROCEDURE bi.process_bi_postclassification()
    OWNER TO useradmin;
