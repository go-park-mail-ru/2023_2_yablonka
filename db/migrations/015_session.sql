CREATE TABLE IF NOT EXISTS public.session
(
    token character varying(64) NOT NULL,
    expiration_date timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP + INTERVAL '14 days',
    id_user serial NOT NULL,
    CONSTRAINT session_pkey PRIMARY KEY (token),
    CONSTRAINT session_id_user_fkey FOREIGN KEY (id_user)
        REFERENCES public."user" (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
        NOT VALID
)

---- create above / drop below ----

DROP TABLE IF EXISTS public.session;
