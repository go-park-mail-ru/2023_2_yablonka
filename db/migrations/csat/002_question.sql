CREATE TABLE IF NOT EXISTS csat_db.question
(
    id integer NOT NULL DEFAULT nextval('csat_db.question_id_seq'::regclass),
    content text COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT question_pkey PRIMARY KEY (id)
)

---- create above / drop below ----

DROP TABLE IF EXISTS csat_db.question;