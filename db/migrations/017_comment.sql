CREATE TABLE IF NOT EXISTS public.comment
(
    id serial NOT NULL,
    id_user serial NOT NULL,
    id_task serial NOT NULL,
    content text NOT NULL,
    date_created timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT comment_pkey PRIMARY KEY (id),
    CONSTRAINT comment_id_task_fkey FOREIGN KEY (id_task)
        REFERENCES public.task (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE CASCADE
        NOT VALID,
    CONSTRAINT comment_id_user_fkey FOREIGN KEY (id_user)
        REFERENCES public."user" (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE SET NULL
        NOT VALID,
    CONSTRAINT content_length_check CHECK (length(content) <= 2000) NOT VALID
)

---- create above / drop below ----

DROP TABLE IF EXISTS public.comment;
