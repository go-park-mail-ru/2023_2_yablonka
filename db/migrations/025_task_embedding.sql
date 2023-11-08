CREATE TABLE IF NOT EXISTS public.task_embedding
(
    id_embedding serial NOT NULL,
    id_task serial NOT NULL,
    CONSTRAINT task_embedding_pkey PRIMARY KEY (id_embedding, id_task),
    CONSTRAINT task_embedding_id_embedding_fkey FOREIGN KEY (id_embedding)
        REFERENCES public.embedding (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
        NOT VALID,
    CONSTRAINT task_embedding_id_task_fkey FOREIGN KEY (id_task)
        REFERENCES public.task (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
        NOT VALID
)

---- create above / drop below ----

DROP TABLE IF EXISTS public.task_embedding;
