CREATE TABLE IF NOT EXISTS public.answer
(
    id serial NOT NULL,
    id_user serial NOT NULL,
    id_question serial NOT NULL,
    score smallint NOT NULL DEFAULT 0,
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