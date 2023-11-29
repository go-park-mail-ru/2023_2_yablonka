
CREATE TABLE IF NOT EXISTS public.favourite_boards
(
    id_board serial NOT NULL,
    id_user serial NOT NULL,
    CONSTRAINT favourite_boards_pkey PRIMARY KEY (id_board, id_user),
    CONSTRAINT favourite_boards_id_board_fkey FOREIGN KEY (id_board)
        REFERENCES public.board (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE CASCADE
        NOT VALID,
    CONSTRAINT favourite_boards_id_user_fkey FOREIGN KEY (id_user)
        REFERENCES public."user" (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE CASCADE
        NOT VALID
)

---- create above / drop below ----

DROP TABLE IF EXISTS public.favourite_boards;