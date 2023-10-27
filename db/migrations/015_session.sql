
CREATE TABLE IF NOT EXISTS public.Session
(
    token character varying(64) NOT NULL,
    expiration_date timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP + INTERVAL '1 day',
    id_user serial NOT NULL,
    PRIMARY KEY (token)
        INCLUDE(token),
    UNIQUE (token, id_user)
        INCLUDE(token, id_user)
);

ALTER TABLE IF EXISTS public.Session
    ADD FOREIGN KEY (id_user)
    REFERENCES public.User (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;


---- create above / drop below ----

DROP TABLE IF EXISTS public.Session;
