CREATE TABLE IF NOT EXISTS public.task_template
(
    id serial NOT NULL,
    data json NOT NULL,
    PRIMARY KEY (id)
);

---- create above / drop below ----

DROP TABLE IF EXISTS public.task_template;
