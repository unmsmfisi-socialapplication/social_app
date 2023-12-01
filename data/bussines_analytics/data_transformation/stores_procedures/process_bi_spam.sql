
--Store Procedure spam


CREATE PROCEDURE process_bi_spam
AS
BEGIN
    -- Paso 1: Vaciar la tabla bi_spam_transformada
    TRUNCATE TABLE bi_spam_transformada;

    -- Paso 2: Transformar y limpiar los datos
    -- Esto incluirá eliminar duplicados, nulos, blancos y caracteres no deseados
    -- También se asegurará de que la columna 'category' tenga solo los valores -1, 0, 1
    INSERT INTO bi_spam_transformada  
    SELECT DISTINCT
        CONVERT(nvarchar(max), FORMATTED_CONTENT) AS FORMATTED_CONTENT,
        CASE WHEN PREDICCION IN ('0', '1') THEN PREDICCION ELSE NULL END,
        GETDATE() AS fecha_sp
    FROM [dbo].[BI_Spam]
    WHERE FORMATTED_CONTENT IS NOT NULL AND LTRIM(RTRIM(CONVERT(VARCHAR(MAX), FORMATTED_CONTENT))) != '';

    -- Eliminar las filas con Post NULL después de la limpieza
    DELETE FROM bi_spam_transformada WHERE FORMATTED_CONTENT IS NULL;
END;
GO

 