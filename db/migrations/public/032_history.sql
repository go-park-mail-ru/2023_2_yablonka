CREATE TABLE IF NOT EXISTS public.edit_history
(
    id serial NOT NULL,
    id_user serial NOT NULL,
    id_board serial NOT NULL,
    edit_date timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP, 
    edit_summary text NOT NULL DEFAULT 'blank',
    CONSTRAINT edit_id_user_fkey FOREIGN KEY (id_user)
        REFERENCES public."user" (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE SET NULL
        NOT VALID,
    CONSTRAINT edit_id_board_fkey FOREIGN KEY (id_board)
        REFERENCES public."board" (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE SET NULL
        NOT VALID
)

---- create above / drop below ----

DROP TABLE IF EXISTS public.edit_history;
