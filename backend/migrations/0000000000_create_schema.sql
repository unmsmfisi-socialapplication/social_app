DO $$
DECLARE
    search_path_value text := current_setting('search_path');
BEGIN

    -- Check if the schema exists
    IF (SELECT COUNT(1) FROM pg_namespace WHERE nspname = search_path_value) = 0 THEN
        -- If the schema does not exist, create it
        EXECUTE 'CREATE SCHEMA ' || (search_path_value);
    END IF;
END $$;