CREATE TABLE IF NOT EXISTS public.task_file
(
    id_task serial NOT NULL,
    id_file serial NOT NULL,
    CONSTRAINT task_file_pkey PRIMARY KEY (id_task, id_file),
    CONSTRAINT task_file_id_file_fkey FOREIGN KEY (id_file)
        REFERENCES public.file (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE CASCADE,
    CONSTRAINT task_file_id_task_fkey FOREIGN KEY (id_task)
        REFERENCES public.task (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE CASCADE
)

---- create above / drop below ----

DROP TABLE IF EXISTS public.task_file;
