CREATE TABLE IF NOT EXISTS public.Board_User
(
    id_board serial NOT NULL,
    id_user serial NOT NULL,
    PRIMARY KEY (id_board, id_user),
    UNIQUE (id_board, id_user)
        INCLUDE(id_board, id_user)
);

ALTER TABLE IF EXISTS public.Board_User
    ADD FOREIGN KEY (id_board)
    REFERENCES public.Board (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;


ALTER TABLE IF EXISTS public.Board_User
    ADD FOREIGN KEY (id_user)
    REFERENCES public.User (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;


---- create above / drop below ----

DROP TABLE IF EXISTS public.Board_User;