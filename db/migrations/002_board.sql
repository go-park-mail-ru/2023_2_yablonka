CREATE TABLE IF NOT EXISTS public.Board
(
    id serial NOT NULL,
    id_workspace serial NOT NULL,
    name character varying(150) NOT NULL DEFAULT 'Доска',
    description text,
    date_created timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    thumbnail_url character varying(2048),
    PRIMARY KEY (id),
    UNIQUE (id)
        INCLUDE(id)
);

ALTER TABLE IF EXISTS public.Board
    ADD FOREIGN KEY (id_workspace)
    REFERENCES public.Workspace (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;


---- create above / drop below ----

DROP TABLE IF EXISTS public.Board;

