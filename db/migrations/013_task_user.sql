CREATE TABLE IF NOT EXISTS public.task_user
(
    id_user bigint NOT NULL,
    id_task bigint NOT NULL,
    CONSTRAINT task_user_pkey PRIMARY KEY (id_user, id_task),
    CONSTRAINT task_user_id_task_fkey FOREIGN KEY (id_task)
        REFERENCES public.task (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE CASCADE
        NOT VALID,
    CONSTRAINT task_user_id_user_fkey FOREIGN KEY (id_user)
        REFERENCES public."user" (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE CASCADE
        NOT VALID
)

---- create above / drop below ----

DROP TABLE IF EXISTS public.task_user;
