CREATE TABLE IF NOT EXISTS public.Board_template
(
    id serial NOT NULL,
    data json NOT NULL,
    PRIMARY KEY (id),
    UNIQUE (id)
        INCLUDE(id)
);

---- create above / drop below ----

DROP TABLE IF EXISTS public.Board_template;
