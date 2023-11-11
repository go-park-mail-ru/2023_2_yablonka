CREATE TABLE IF NOT EXISTS public.csrf
(
    token text NOT NULL,
    id_user serial NOT NULL,
    expiration_date timestamp without time zone NOT NULL,
    CONSTRAINT csrf_pkey PRIMARY KEY (token),
    CONSTRAINT csrf_id_user_fkey FOREIGN KEY (id_user)
        REFERENCES public."user" (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
)

---- create above / drop below ----

DROP TABLE IF EXISTS public.csrf;
