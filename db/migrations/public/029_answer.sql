CREATE TABLE IF NOT EXISTS public.answer
(
    id integer NOT NULL DEFAULT nextval('public.response_id_seq'::regclass),
    id_user integer NOT NULL DEFAULT nextval('public.response_id_user_seq'::regclass),
    score smallint NOT NULL DEFAULT 0,
    id_question integer NOT NULL DEFAULT nextval('public.response_id_question_seq'::regclass),
    CONSTRAINT response_pkey PRIMARY KEY (id),
    CONSTRAINT answer_id_question_fkey FOREIGN KEY (id_question)
        REFERENCES public.question (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
        NOT VALID,
    CONSTRAINT response_id_user_fkey FOREIGN KEY (id_user)
        REFERENCES public."user" (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
)

---- create above / drop below ----

DROP TABLE IF EXISTS public.answer;