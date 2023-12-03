CREATE TABLE IF NOT EXISTS public.comment_reply
(
    id_reply serial NOT NULL,
    id_comment serial NOT NULL,
    CONSTRAINT comment_reply_pkey PRIMARY KEY (id_reply),
    CONSTRAINT original_comment FOREIGN KEY (id_comment)
        REFERENCES public.comment (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE SET NULL
        NOT VALID,
    CONSTRAINT reply FOREIGN KEY (id_reply)
        REFERENCES public.comment (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE SET NULL
        NOT VALID
)

---- create above / drop below ----

DROP TABLE IF EXISTS public.comment_reply;
