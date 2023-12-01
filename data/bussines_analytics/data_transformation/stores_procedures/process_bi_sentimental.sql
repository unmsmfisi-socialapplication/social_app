

--Store Procedure sentimental

CREATE OR ALTER PROCEDURE process_bi_sentimental
AS
BEGIN
    -- Paso 1: Vaciar la tabla bi_sentimental_transformada
    TRUNCATE TABLE dbo.bi_sentimental_transformada;

    -- Paso 2: Transformar y limpiar los datos
    -- Esto incluirá eliminar duplicados, nulos, blancos y caracteres no deseados
    -- También se asegurará de que la columna 'category' tenga solo los valores -1, 0, 1
    INSERT INTO dbo.bi_sentimental_transformada (clean_text, category, fecha_actualizacion)
    
        SELECT DISTINCT
            CONVERT(nvarchar(max), clean_text) AS clean_text,
            CASE WHEN category IN (-1, 0, 1) THEN category ELSE NULL END AS category,
            GETDATE() AS fecha_actualizacion
        FROM dbo.BI_Sentiment
        WHERE clean_text IS NOT NULL AND LTRIM(RTRIM(CONVERT(nvarchar(max), clean_text))) != '';

    -- Eliminar las filas con category NULL después de la limpieza
    DELETE FROM dbo.bi_sentimental_transformada WHERE category IS NULL;
END; 