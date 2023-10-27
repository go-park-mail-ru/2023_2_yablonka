CREATE TABLE IF NOT EXISTS public.Tag
(
    id serial NOT NULL,
    name character varying(35) NOT NULL,
    color character varying(6) NOT NULL DEFAULT 'FFFFFF',
    PRIMARY KEY (id),
    UNIQUE (name, id)
        INCLUDE(name, id)
);

---- create above / drop below ----

DROP TABLE IF EXISTS public.Tag;
