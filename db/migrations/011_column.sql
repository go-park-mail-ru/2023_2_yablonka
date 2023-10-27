CREATE TABLE IF NOT EXISTS public.Column
(
    id serial NOT NULL,
    id_board serial NOT NULL,
    name character varying(150) NOT NULL DEFAULT 'Столбец',
    description text,
    list_position smallint NOT NULL,
    PRIMARY KEY (id)
        INCLUDE(id),
    UNIQUE (id)
        INCLUDE(id)
);

ALTER TABLE IF EXISTS public.Column
    ADD FOREIGN KEY (id_board)
    REFERENCES public.Board (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;


---- create above / drop below ----

DROP TABLE IF EXISTS public.Column;
