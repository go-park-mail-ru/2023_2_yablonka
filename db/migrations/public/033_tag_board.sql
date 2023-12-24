CREATE TABLE IF NOT EXISTS public.tag_board
(
    id_tag serial NOT NULL,
    id_board serial NOT NULL,
    CONSTRAINT tag_board_pkey PRIMARY KEY (id_tag, id_board),
    CONSTRAINT tag_board_id_tag_fkey FOREIGN KEY (id_tag)
        REFERENCES public.tag (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE CASCADE
        NOT VALID,
    CONSTRAINT tag_board_id_board_fkey FOREIGN KEY (id_board)
        REFERENCES public.board (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE CASCADE
        NOT VALID
)

---- create above / drop below ----

DROP TABLE IF EXISTS public.tag_board;
