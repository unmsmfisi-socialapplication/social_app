--Testing to SP, Tiempo de ejecución

DO $$
DECLARE
    startTime TIMESTAMP;
    endTime TIMESTAMP;
BEGIN
    startTime := clock_timestamp();
    CALL bi.process_bi_sentimental();
    endTime := clock_timestamp();

    RAISE NOTICE 'Tiempo de ejecución: %', (endTime - startTime);
END;
$$;

-- Cuanta data se ha eliminado según lo configurado, data eliminada 

SELECT sentimental,transformada,(sentimental-transformada) as eliminados 
FROM (
	SELECT COUNT(*) as sentimental FROM bi.bi_sentimental) as  s
	,(SELECT COUNT(*) as transformada FROM bi.bi_sentimental_transformada) as t;

-- ¿Por qué se eliminaron?
	--Nulos o en blanco
SELECT COUNT(*) as "Eliminados Por Nulos o en Blanco"
	FROM bi.bi_sentimental WHERE clean_text IS NULL OR TRIM(clean_text) = '';
 -- resultados en categoria distintos a -1 0 1
SELECT COUNT(*) "Eliminados resultados en categoria distintos @ lo requerido"
	FROM bi.bi_sentimental WHERE category NOT IN ('-1','0','1');




