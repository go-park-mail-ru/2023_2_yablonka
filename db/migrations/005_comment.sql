CREATE TABLE IF NOT EXISTS public.Comment
(
    id serial NOT NULL,
    id_user serial NOT NULL,
    id_task serial NOT NULL,
    content text NOT NULL,
    date_created timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id_user),
    UNIQUE (id)
        INCLUDE(id)
);

ALTER TABLE IF EXISTS public.Comment
    ADD FOREIGN KEY (id_user)
    REFERENCES public.User (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;


ALTER TABLE IF EXISTS public.Comment
    ADD FOREIGN KEY (id_task)
    REFERENCES public.Task (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;


---- create above / drop below ----

DROP TABLE IF EXISTS public.Comment;
