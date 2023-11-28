CREATE OR REPLACE PROCEDURE bi.process_bi_sentimental()
LANGUAGE plpgsql
AS $$
BEGIN
    -- Paso 1: Vaciar la tabla bi_sentimental_transformada
    TRUNCATE TABLE bi.bi_sentimental_transformada;
    
    -- Paso 2: Transformar y limpiar los datos
    -- Esto incluirá eliminar duplicados, nulos, blancos y caracteres no deseados
    -- También se asegurará de que la columna 'category' tenga solo los valores -1, 0, 1
    INSERT INTO bi.bi_sentimental_transformada (clean_text, category, fecha_actualizacion)
    SELECT DISTINCT
        TRIM(REGEXP_REPLACE(clean_text, '[^a-zA-Z0-9ÁÉÍÓÚáéíóúñÑ\s]', '', 'g')),
        CASE WHEN category IN ('-1', '0', '1') THEN category ELSE NULL END,
        CURRENT_TIMESTAMP
    FROM bi.bi_sentimental
    WHERE clean_text IS NOT NULL AND TRIM(clean_text) != '';

    -- Eliminar las filas con category NULL después de la limpieza
    DELETE FROM bi.bi_sentimental_transformada WHERE category IS NULL;
END;
$$;


