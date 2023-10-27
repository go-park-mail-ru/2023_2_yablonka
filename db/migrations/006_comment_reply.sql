CREATE TABLE IF NOT EXISTS public.Comment_reply
(
    id_reply serial NOT NULL,
    id_comment serial NOT NULL,
    PRIMARY KEY (id_reply)
        INCLUDE(id_reply),
    UNIQUE (id_reply)
);

ALTER TABLE IF EXISTS public.Comment_reply
    ADD CONSTRAINT original_comment FOREIGN KEY (id_comment)
    REFERENCES public.Comment (id) MATCH SIMPLE
    ON UPDATE CASCADE
    ON DELETE CASCADE
    NOT VALID;


ALTER TABLE IF EXISTS public.Comment_reply
    ADD CONSTRAINT reply FOREIGN KEY (id_reply)
    REFERENCES public.Comment (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;


---- create above / drop below ----

DROP TABLE IF EXISTS public.Comment_reply;
