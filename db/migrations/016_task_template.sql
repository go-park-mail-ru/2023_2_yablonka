CREATE TABLE IF NOT EXISTS public.Task_template
(
    id serial NOT NULL,
    data json NOT NULL,
    PRIMARY KEY (id)
        INCLUDE(id),
    UNIQUE (id)
);

---- create above / drop below ----

DROP TABLE IF EXISTS public.Task_template;
