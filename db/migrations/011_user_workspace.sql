CREATE TABLE IF NOT EXISTS public.User_Workspace
(
    id_user serial NOT NULL,
    id_workspace serial NOT NULL,
    id_role serial NOT NULL,
    PRIMARY KEY (id_user, id_workspace)
        INCLUDE(id_user, id_workspace),
    UNIQUE (id_user, id_workspace)
        INCLUDE(id_user, id_workspace)
);

ALTER TABLE IF EXISTS public.User_Workspace
    ADD FOREIGN KEY (id_user)
    REFERENCES public.User (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;


ALTER TABLE IF EXISTS public.User_Workspace
    ADD FOREIGN KEY (id_role)
    REFERENCES public.Role (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;


ALTER TABLE IF EXISTS public.User_Workspace
    ADD FOREIGN KEY (id_workspace)
    REFERENCES public.Workspace (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;


---- create above / drop below ----

DROP TABLE IF EXISTS public.User_Workspace;
