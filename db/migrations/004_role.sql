CREATE TABLE IF NOT EXISTS public.Role
(
    id serial NOT NULL,
    name character varying(100) NOT NULL DEFAULT 'Роль',
    description text,
    CONSTRAINT pk_role PRIMARY KEY (id),
    UNIQUE (id)
        INCLUDE(id)
);

---- create above / drop below ----

DROP TABLE IF EXISTS public.Role;
