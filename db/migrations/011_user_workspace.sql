CREATE TABLE IF NOT EXISTS public.user_workspace
(
    id_user serial NOT NULL,
    id_workspace serial NOT NULL,
    id_role serial NOT NULL,
    CONSTRAINT user_workspace_pkey PRIMARY KEY (id_user, id_workspace),
    CONSTRAINT user_workspace_id_role_fkey FOREIGN KEY (id_role)
        REFERENCES public.role (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
        NOT VALID,
    CONSTRAINT user_workspace_id_user_fkey FOREIGN KEY (id_user)
        REFERENCES public."user" (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
        NOT VALID,
    CONSTRAINT user_workspace_id_workspace_fkey FOREIGN KEY (id_workspace)
        REFERENCES public.workspace (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
        NOT VALID
)

---- create above / drop below ----

DROP TABLE IF EXISTS public.user_workspace;
