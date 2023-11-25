CREATE TABLE IF NOT EXISTS public.board_template
(
    id serial NOT NULL,
    data json NOT NULL,
    PRIMARY KEY (id)
);

---- create above / drop below ----

DROP TABLE IF EXISTS public.board_template;
