CREATE TABLE IF NOT EXISTS public.workspace
(
    id serial NOT NULL,
    name text NOT NULL DEFAULT 'Рабочее место',
    date_created timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    description text,
    CONSTRAINT workspace_pkey PRIMARY KEY (id),
    CONSTRAINT workspace_name_length_check CHECK (length(name) <= 150) NOT VALID
)

---- create above / drop below ----

DROP TABLE IF EXISTS public.workspace;
