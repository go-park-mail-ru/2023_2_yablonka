CREATE TABLE IF NOT EXISTS public.workspace
(
    id serial NOT NULL,
    name character varying(150) NOT NULL DEFAULT 'Рабочее место',
    thumbnail_url text,
    date_created timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    description text,
    CONSTRAINT workspace_pkey PRIMARY KEY (id)
)

---- create above / drop below ----

DROP TABLE IF EXISTS public.workspace;
