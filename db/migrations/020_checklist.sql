CREATE TABLE IF NOT EXISTS public.Checklist
(
    id serial NOT NULL,
    id_task serial NOT NULL,
    name character varying(100) NOT NULL DEFAULT 'Чек-лист',
    list_position smallint NOT NULL,
    PRIMARY KEY (id),
    UNIQUE (id)
        INCLUDE(id)
);

ALTER TABLE IF EXISTS public.Checklist
    ADD FOREIGN KEY (id_task)
    REFERENCES public.Task (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;


---- create above / drop below ----

DROP TABLE IF EXISTS public.Checklist;
