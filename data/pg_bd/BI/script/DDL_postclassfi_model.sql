CREATE TABLE IF NOT EXISTS bi.ResultsPostClasif
(
    "ID" INT UNIQUE,
    post_description TEXT ,
    category TEXT ,
    model_version VARCHAR[20] ,
    feh_create TIMESTAMP ,
	name_user VARCHAR[20]
)