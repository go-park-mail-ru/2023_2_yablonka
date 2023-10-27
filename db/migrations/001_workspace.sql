CREATE TABLE IF NOT EXISTS public.Workspace
(
    id serial NOT NULL,
    name character varying(150) NOT NULL DEFAULT 'Рабочее место',
    thumbnail_url character varying(2048),
    date_created timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    description text,
    PRIMARY KEY (id),
    UNIQUE (id)
        INCLUDE(id)
);

---- create above / drop below ----

DROP TABLE IF EXISTS public.Workspace;
