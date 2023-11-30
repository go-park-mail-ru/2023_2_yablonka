CREATE TABLE IF NOT EXISTS public.comment_embedding
(
    id_embedding serial NOT NULL,
    id_comment serial NOT NULL,
    CONSTRAINT comment_embedding_pkey PRIMARY KEY (id_embedding, id_comment),
    CONSTRAINT comment_embedding_id_embedding_fkey FOREIGN KEY (id_embedding)
        REFERENCES public.embedding (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE CASCADE
        NOT VALID,
    CONSTRAINT comment_embedding_id_comment_fkey FOREIGN KEY (id_comment)
        REFERENCES public.comment (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE CASCADE
        NOT VALID
)

---- create above / drop below ----

DROP TABLE IF EXISTS public.comment_embedding;