CREATE TABLE IF NOT EXISTS public.task
(
    id serial NOT NULL,
    id_list serial NOT NULL,
    date_created timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    description text,
    name text NOT NULL DEFAULT 'Задача',
    "task_start" timestamp without time zone,
    "task_end" timestamp without time zone,
    list_position smallint NOT NULL,
    CONSTRAINT task_pkey PRIMARY KEY (id),
    CONSTRAINT task_id_list_fkey FOREIGN KEY (id_list)
        REFERENCES public.list (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE CASCADE
        NOT VALID,
    CONSTRAINT task_list_position_check CHECK (list_position >= 0) NOT VALID,
    CONSTRAINT task_name_length_check CHECK (length(name) <= 150) NOT VALID
)

---- create above / drop below ----

DROP TABLE IF EXISTS public.task;
