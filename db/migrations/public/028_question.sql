CREATE TABLE IF NOT EXISTS public.question
(
    id integer NOT NULL DEFAULT nextval('public.question_id_seq'::regclass),
    content text COLLATE pg_catalog."default" NOT NULL,
    id_type integer NOT NULL DEFAULT nextval('public.question_id_type_seq'::regclass),
    CONSTRAINT question_pkey PRIMARY KEY (id),
    CONSTRAINT question_id_type_fkey FOREIGN KEY (id_type)
        REFERENCES public.question_type (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
        NOT VALID
)

---- create above / drop below ----

DROP TABLE IF EXISTS public.question;