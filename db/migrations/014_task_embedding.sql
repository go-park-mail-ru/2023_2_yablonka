CREATE TABLE IF NOT EXISTS public.Task_Embedding
(
    id serial NOT NULL,
    id_task serial NOT NULL,
    id_user serial NOT NULL,
    url character varying(2048) NOT NULL,
    PRIMARY KEY (id),
    UNIQUE (id)
);

ALTER TABLE IF EXISTS public.Task_Embedding
    ADD FOREIGN KEY (id_task)
    REFERENCES public.Task (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;


ALTER TABLE IF EXISTS public.Task_Embedding
    ADD FOREIGN KEY (id_user)
    REFERENCES public.User (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;


---- create above / drop below ----

DROP TABLE IF EXISTS public.Task_Embedding;
