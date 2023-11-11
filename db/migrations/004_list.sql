CREATE TABLE IF NOT EXISTS public.list
(
    id serial NOT NULL,
    id_board serial NOT NULL,
    name text NOT NULL DEFAULT 'Столбец',
    description text,
    list_position smallint NOT NULL,
    CONSTRAINT list_pkey PRIMARY KEY (id),
    CONSTRAINT list_id_board_fkey FOREIGN KEY (id_board)
        REFERENCES public.board (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE CASCADE
        NOT VALID,
    CONSTRAINT list_list_position_check CHECK (list_position >= 0) NOT VALID,
    CONSTRAINT column_name_length_check CHECK (length(name) <= 150) NOT VALID
)

---- create above / drop below ----

DROP TABLE IF EXISTS public.list;
