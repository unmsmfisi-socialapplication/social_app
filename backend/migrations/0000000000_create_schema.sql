DO $$
DECLARE
    search_path_value text;
BEGIN
    SELECT current_setting('search_path') INTO search_path_value;

    -- Check if the schema exists
    IF NOT EXISTS (SELECT 1 FROM pg_namespace WHERE nspname = search_path_value) THEN
        -- If the schema does not exist, create it
        EXECUTE 'CREATE SCHEMA ' || quote_ident(search_path_value);
    END IF;
END $$;