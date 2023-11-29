CREATE TABLE IF NOT EXISTS public.board_user
(
    id_board serial NOT NULL,
    id_user serial NOT NULL,
    CONSTRAINT board_user_pkey PRIMARY KEY (id_board, id_user),
    CONSTRAINT board_user_id_board_fkey FOREIGN KEY (id_board)
        REFERENCES public.board (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE CASCADE
        NOT VALID,
    CONSTRAINT board_user_id_user_fkey FOREIGN KEY (id_user)
        REFERENCES public."user" (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE CASCADE
        NOT VALID
)

---- create above / drop below ----

DROP TABLE IF EXISTS public.board_user;