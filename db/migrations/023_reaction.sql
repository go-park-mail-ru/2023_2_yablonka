CREATE TABLE IF NOT EXISTS public.Reaction
(
    id serial NOT NULL,
    id_comment serial NOT NULL,
    id_user serial NOT NULL,
    content character varying(2) NOT NULL,
    PRIMARY KEY (id)
        INCLUDE(id),
    UNIQUE (id)
        INCLUDE(id)
);

ALTER TABLE IF EXISTS public.Reaction
    ADD FOREIGN KEY (id_user)
    REFERENCES public.User (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;


ALTER TABLE IF EXISTS public.Reaction
    ADD FOREIGN KEY (id_comment)
    REFERENCES public.Comment (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;


---- create above / drop below ----

DROP TABLE IF EXISTS public.Reaction;
