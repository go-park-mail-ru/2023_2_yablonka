CREATE TABLE IF NOT EXISTS public.Task_User
(
    id_user bigint NOT NULL,
    id_task bigint NOT NULL,
    PRIMARY KEY (id_user, id_task)
        INCLUDE(id_user, id_task),
    UNIQUE (id_user, id_task)
        INCLUDE(id_user, id_task)
);

ALTER TABLE IF EXISTS public.Task_User
    ADD FOREIGN KEY (id_task)
    REFERENCES public.Task (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;


ALTER TABLE IF EXISTS public.Task_User
    ADD FOREIGN KEY (id_user)
    REFERENCES public."User" (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;


---- create above / drop below ----

DROP TABLE IF EXISTS public.Task_User;
