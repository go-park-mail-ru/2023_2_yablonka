CREATE TABLE IF NOT EXISTS public.task
(
    id serial NOT NULL,
    id_list serial NOT NULL,
    date_created timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    description text,
    name character varying(150) NOT NULL DEFAULT 'Задача',
    "start" timestamp without time zone,
    "end" timestamp without time zone,
    list_postition smallint NOT NULL,
    CONSTRAINT task_pkey PRIMARY KEY (id),
    CONSTRAINT task_id_list_fkey FOREIGN KEY (id_list)
        REFERENCES public.list (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
        NOT VALID,
    CONSTRAINT task_list_postition_check CHECK (list_postition >= 0) NOT VALID
)

---- create above / drop below ----

DROP TABLE IF EXISTS public.task;
