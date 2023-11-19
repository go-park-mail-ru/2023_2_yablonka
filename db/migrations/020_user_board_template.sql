CREATE TABLE IF NOT EXISTS public.user_board_template
(
    id_user serial NOT NULL,
    id_template serial NOT NULL,
    CONSTRAINT user_board_template_pkey PRIMARY KEY (id_user, id_template),
    CONSTRAINT user_board_template_id_template_fkey FOREIGN KEY (id_template)
        REFERENCES public.board_template (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE CASCADE
        NOT VALID,
    CONSTRAINT user_board_template_id_user_fkey FOREIGN KEY (id_user)
        REFERENCES public."user" (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE CASCADE
        NOT VALID
)

---- create above / drop below ----

DROP TABLE IF EXISTS public.user_board_template;
