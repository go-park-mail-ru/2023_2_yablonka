CREATE TABLE IF NOT EXISTS public.embedding
(
    id serial NOT NULL,
    id_user serial NOT NULL,
    url text NOT NULL,
    CONSTRAINT embedding_pkey PRIMARY KEY (id),
    CONSTRAINT embedding_id_user_fkey FOREIGN KEY (id_user)
        REFERENCES public."user" (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE SET NULL
        NOT VALID
)

---- create above / drop below ----

DROP TABLE IF EXISTS public.embedding;
