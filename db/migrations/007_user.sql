CREATE TABLE IF NOT EXISTS public.User
(
    id serial NOT NULL,
    email character varying(256) NOT NULL,
    password_hash character varying(256) NOT NULL,
    name character varying(100),
    surname character varying(100),
    avatar_url character varying(2048),
    description text,
    PRIMARY KEY (id)
        INCLUDE(id),
    UNIQUE (email, id)
        INCLUDE(email, id)
);

---- create above / drop below ----

DROP TABLE IF EXISTS public.User;
