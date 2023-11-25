CREATE TABLE IF NOT EXISTS csat_db.question_type
(
    id integer NOT NULL DEFAULT nextval('csat_db.question_type_id_seq'::regclass),
    name text COLLATE pg_catalog."default" NOT NULL,
    max_score smallint NOT NULL,
    CONSTRAINT question_type_pkey PRIMARY KEY (id),
    CONSTRAINT question_type_name_key UNIQUE (name)
)

---- create above / drop below ----

DROP TABLE IF EXISTS csat_db.question_type;
