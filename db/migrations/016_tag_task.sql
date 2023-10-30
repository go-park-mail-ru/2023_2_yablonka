CREATE TABLE IF NOT EXISTS public.tag_task
(
    id_tag serial NOT NULL,
    id_task serial NOT NULL,
    CONSTRAINT tag_task_pkey PRIMARY KEY (id_tag, id_task),
    CONSTRAINT tag_task_id_tag_fkey FOREIGN KEY (id_tag)
        REFERENCES public.tag (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
        NOT VALID,
    CONSTRAINT tag_task_id_task_fkey FOREIGN KEY (id_task)
        REFERENCES public.task (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
        NOT VALID
)

---- create above / drop below ----

DROP TABLE IF EXISTS public.tag_task;
