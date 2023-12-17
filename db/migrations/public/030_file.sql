CREATE TABLE IF NOT EXISTS public.file
(
    id serial NOT NULL,
    name text NOT NULL DEFAULT 'blank',
    filepath text NOT NULL,
    date_created timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT files_pkey PRIMARY KEY (id)
)

---- create above / drop below ----

DROP TABLE IF EXISTS public.file;