CREATE TABLE IF NOT EXISTS public.Task
(
    id serial NOT NULL,
    id_column serial NOT NULL,
    date_created timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    description text,
    name character varying(150) NOT NULL DEFAULT 'Задача',
    start timestamp without time zone,
    "end" timestamp without time zone,
    list_postition smallint NOT NULL,
    PRIMARY KEY (id),
    UNIQUE (id)
        INCLUDE(id)
);

ALTER TABLE IF EXISTS public.Task
    ADD FOREIGN KEY (id_column)
    REFERENCES public."Column" (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;


---- create above / drop below ----

DROP TABLE IF EXISTS public.Task;
