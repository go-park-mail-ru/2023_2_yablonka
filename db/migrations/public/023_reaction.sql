CREATE TABLE IF NOT EXISTS public.reaction
(
    id serial NOT NULL,
    id_comment serial NOT NULL,
    id_user serial NOT NULL,
    content text NOT NULL,
    CONSTRAINT reaction_pkey PRIMARY KEY (id),
    CONSTRAINT reaction_id_comment_fkey FOREIGN KEY (id_comment)
        REFERENCES public.comment (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE CASCADE
        NOT VALID,
    CONSTRAINT reaction_id_user_fkey FOREIGN KEY (id_user)
        REFERENCES public."user" (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE SET NULL
        NOT VALID,
    CONSTRAINT reaction_content_length_check CHECK (length(content) <= 2) NOT VALID
)

---- create above / drop below ----

DROP TABLE IF EXISTS public.reaction;
