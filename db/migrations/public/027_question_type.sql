BEGIN
CREATE TABLE IF NOT EXISTS public.question_type
(
    id serial NOT NULL,
    name text COLLATE pg_catalog."default" NOT NULL,
    max_score smallint NOT NULL,
    CONSTRAINT question_type_pkey PRIMARY KEY (id),
    CONSTRAINT question_type_name_key UNIQUE (name)
);

INSERT INTO public.question_type
    (name, max_score)
	VALUES 
    ('NSP', 5),
	('CSI', 10);

END;

---- create above / drop below ----

DROP TABLE IF EXISTS public.question_type;
