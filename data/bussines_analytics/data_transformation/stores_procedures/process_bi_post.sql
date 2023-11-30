

--Store Procedure post

CREATE PROCEDURE process_bi_post
AS
BEGIN
    -- Paso 1: Vaciar la tabla bi_post_transformada
    TRUNCATE TABLE bi_post_transformada;

    -- Paso 2: Transformar y limpiar los datos
    -- Esto incluirá eliminar duplicados, nulos, blancos y caracteres no deseados
    -- También se asegurará de que la columna 'category' tenga solo los valores -1, 0, 1
    INSERT INTO bi_post_transformada (clean_text, category, fecha_actualizacion)
    SELECT DISTINCT
        CONVERT(nvarchar(max), clean_text) AS clean_text,
        CASE 
			WHEN category_label LIKE '%ENTERTAIN%' THEN 'ENTERTAINMENT'
			WHEN category_label LIKE '%POLITI%' THEN 'POLITICS'
			WHEN category_label LIKE '%WORLD%' THEN 'WORLD NEWS'
			WHEN category_label LIKE '%WEIRD%' THEN 'WEIRD NEWS'
			WHEN category_label LIKE '%ARTS & CULTURE%' THEN 'ARTS & CULTURE'
			WHEN category_label LIKE '%MED%' THEN 'COMEDY'
			WHEN category_label LIKE '%SCIENCE%' THEN 'SCIENCE & TECH'
			WHEN category_label LIKE '%EDUCAT%' THEN 'EDUCATION'
			WHEN category_label LIKE '%PARENT%' THEN 'PARENTING'
			WHEN category_label LIKE '%SPOR%' THEN 'SPORTS'
			WHEN category_label LIKE '%ENVIRON%' THEN 'ENVIRONMENT'
			WHEN category_label LIKE '%CRIM%' THEN 'CRIME'
			WHEN category_label LIKE '%WELLN%' THEN 'WELLNESS'
			WHEN category_label LIKE '%BUSINESS%' THEN 'BUSINESS & FINANCES'
			WHEN category_label LIKE '%STYLE%' THEN 'STYLE & BEAUTY'
			WHEN category_label LIKE '%FOOD%' THEN 'FOOD & DRINK'
			WHEN category_label LIKE '%GROUPS%' THEN 'GROUPS VOICES'
			WHEN category_label LIKE '%HOME%' THEN 'HOME & LIVING'
			WHEN category_label LIKE '%WOM%' THEN 'WOMEN'
			WHEN category_label LIKE '%TRAVE%' THEN 'TRAVEL'
			WHEN category_label LIKE '%RELIG%' THEN 'RELIGION'
			WHEN category_label LIKE '%IMPAC%' THEN 'IMPACT'
			WHEN category_label LIKE '%WEDDIN%' THEN 'WEDDINGS'
			WHEN category_label LIKE '%MISCEL%' THEN 'MISCELLANEOUS'
			WHEN category_label LIKE '%DIVOR%' THEN 'DIVORCE'
		ELSE NULL END AS category,
        GETDATE() AS fecha_actualizacion
    FROM  [dbo].[BI_Post]
    WHERE clean_text IS NOT NULL AND LTRIM(RTRIM(CONVERT(nvarchar(max), clean_text))) != '';

    -- Eliminar las filas con category NULL después de la limpieza
    DELETE FROM bi_post_transformada WHERE category IS NULL;
END;
GO
