CREATE TABLE IF NOT EXISTS public."user"
(
    id serial NOT NULL,
    email text NOT NULL,
    password_hash text NOT NULL,
    name text,
    surname text,
    avatar_url text NOT NULL DEFAULT 'img/user_avatars/avatar.jpg',
    description text,
    CONSTRAINT user_pkey PRIMARY KEY (id),
    CONSTRAINT user_email_key UNIQUE (email),
    CONSTRAINT user_email_length_check CHECK (length(email) <= 256) NOT VALID,
    CONSTRAINT user_password_hash_length_check CHECK (length(password_hash) <= 256) NOT VALID,
    CONSTRAINT user_name_length_check CHECK (length(name) <= 100) NOT VALID,
    CONSTRAINT user_surname_length_check CHECK (length(surname) <= 100) NOT VALID
)

---- create above / drop below ----

DROP TABLE IF EXISTS public."user";
