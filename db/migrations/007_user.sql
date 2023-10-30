CREATE TABLE IF NOT EXISTS public."user"
(
    id serial NOT NULL,
    email character varying(256) NOT NULL,
    password_hash character varying(256) NOT NULL,
    name character varying(100),
    surname character varying(100),
    avatar_url text,
    description text,
    CONSTRAINT user_pkey PRIMARY KEY (id),
    CONSTRAINT user_email_key UNIQUE (email)
)

---- create above / drop below ----

DROP TABLE IF EXISTS public."user";
