CREATE TYPE public.roles_enum AS ENUM
    ('reader', 'commenter', 'editor', 'creator');

---- create above / drop below ----

DROP TYPE IF EXISTS public.roles_enum;
