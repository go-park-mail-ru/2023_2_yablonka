CREATE TABLE IF NOT EXISTS public.session
(
    id_session text NOT NULL,
    expiration_date timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP + INTERVAL '14 days',
    id_user serial NOT NULL,
    CONSTRAINT session_pkey PRIMARY KEY (id_session),
    CONSTRAINT session_id_user_fkey FOREIGN KEY (id_user)
        REFERENCES public."user" (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE CASCADE
        NOT VALID,
    CONSTRAINT session_id_session_length_check CHECK (length(id_session) <= 64) NOT VALID
)

---- create above / drop below ----

DROP TABLE IF EXISTS public.session;
