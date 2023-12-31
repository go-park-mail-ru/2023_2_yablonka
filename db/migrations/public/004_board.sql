CREATE TABLE IF NOT EXISTS public.board
(
    id serial NOT NULL,
    id_workspace serial NOT NULL,
    name text NOT NULL DEFAULT 'Доска',
    description text,
    date_created timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    thumbnail_url text,
    CONSTRAINT board_pkey PRIMARY KEY (id),
    CONSTRAINT board_id_workspace_fkey FOREIGN KEY (id_workspace)
        REFERENCES public.workspace (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE CASCADE
        NOT VALID,
    CONSTRAINT board_name_length_check CHECK (length(name) <= 150) NOT VALID
)

---- create above / drop below ----

DROP TABLE IF EXISTS public.board;

