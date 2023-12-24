CREATE TABLE IF NOT EXISTS public.tag
(
    id serial NOT NULL,
    name text NOT NULL,
    color text NOT NULL DEFAULT 'FFFFFF',
    CONSTRAINT tag_pkey PRIMARY KEY (id),
    CONSTRAINT tag_name_id_name1_id1_key UNIQUE (name),
    CONSTRAINT tag_color_length_check CHECK (length(color) <= 6) NOT VALID,
    CONSTRAINT tag_name_length_check CHECK (length(name) <= 35) NOT VALID
)

---- create above / drop below ----

DROP TABLE IF EXISTS public.tag;
