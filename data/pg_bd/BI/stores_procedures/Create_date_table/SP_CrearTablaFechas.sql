CREATE OR REPLACE PROCEDURE bi.create_date_table(
	)
LANGUAGE 'plpgsql'
AS $BODY$
DECLARE
fecha_actual TIMESTAMP := '2023-01-01';
BEGIN
    -- Borrar la tabla si ya existe
    IF EXISTS (SELECT 1 FROM information_schema.tables WHERE table_name = 'date_table') THEN
        DROP TABLE date_table;
    END IF;

    -- Crear la tabla
    CREATE TABLE bi.date_table (
        fecha DATE,
        anio INT,
        mes INT,
        dia INT
    );

    -- Bucle para insertar todas las fechas del año 2023 en la tabla
    WHILE fecha_actual <= '2023-12-31' LOOP
        INSERT INTO bi.date_table (fecha, anio, mes, dia)
        VALUES (fecha_actual, EXTRACT(YEAR FROM fecha_actual), EXTRACT(MONTH FROM fecha_actual), EXTRACT(DAY FROM fecha_actual));

        -- Avanzar a la siguiente fecha
        fecha_actual := fecha_actual + INTERVAL '1 day';
    END LOOP;
END;
$BODY$;
ALTER PROCEDURE bi.create_date_table()
    OWNER TO useradmin;s
