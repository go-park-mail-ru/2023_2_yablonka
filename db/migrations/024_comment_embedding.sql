CREATE TABLE IF NOT EXISTS public.Comment_embedding
(
    id serial NOT NULL,
    id_user serial NOT NULL,
    id_comment serial NOT NULL,
    url character varying(2048) NOT NULL,
    PRIMARY KEY (id)
        INCLUDE(id)
);

ALTER TABLE IF EXISTS public.Comment_embedding
    ADD FOREIGN KEY (id_comment)
    REFERENCES public.Comment (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;


ALTER TABLE IF EXISTS public.Comment_embedding
    ADD FOREIGN KEY (id_user)
    REFERENCES public.User (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;


---- create above / drop below ----

DROP TABLE IF EXISTS public.Comment_embedding;