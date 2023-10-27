CREATE TABLE IF NOT EXISTS public.User_Board_template
(
    id_user serial NOT NULL,
    id_template serial NOT NULL,
    PRIMARY KEY (id_user, id_template),
    UNIQUE (id_user, id_template)
        INCLUDE(id_user, id_template)
);

ALTER TABLE IF EXISTS public.User_Board_template
    ADD FOREIGN KEY (id_user)
    REFERENCES public.User (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;


ALTER TABLE IF EXISTS public.User_Board_template
    ADD FOREIGN KEY (id_template)
    REFERENCES public.Board_template (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;


---- create above / drop below ----

DROP TABLE IF EXISTS public.User_Board_template;
