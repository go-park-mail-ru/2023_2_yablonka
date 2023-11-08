CREATE TABLE IF NOT EXISTS public.checklist
(
    id serial NOT NULL,
    id_task serial NOT NULL,
    name text NOT NULL DEFAULT 'Чек-лист',
    list_position smallint NOT NULL,
    CONSTRAINT checklist_pkey PRIMARY KEY (id),
    CONSTRAINT checklist_id_task_fkey FOREIGN KEY (id_task)
        REFERENCES public.task (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
        NOT VALID,
    CONSTRAINT checklist_list_position_check CHECK (list_position >= 0) NOT VALID,
    CONSTRAINT checklist_name_length_check CHECK (length(name) <= 100) NOT VALID
)

---- create above / drop below ----

DROP TABLE IF EXISTS public.checklist;
