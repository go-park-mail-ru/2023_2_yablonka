CREATE TABLE IF NOT EXISTS public.role
(
    id serial NOT NULL,
    name text NOT NULL DEFAULT 'Роль',
    description text,
    CONSTRAINT pk_role PRIMARY KEY (id)
    CONSTRAINT role_name_length_check CHECK (length(name) <= 100) NOT VALID
)

---- create above / drop below ----

DROP TABLE IF EXISTS public.role;
