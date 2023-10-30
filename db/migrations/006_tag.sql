CREATE TABLE IF NOT EXISTS public.tag
(
    id serial NOT NULL,
    name character varying(35) NOT NULL,
    color character varying(6) NOT NULL DEFAULT 'FFFFFF',
    CONSTRAINT tag_pkey PRIMARY KEY (id),
    CONSTRAINT tag_name_id_name1_id1_key UNIQUE (name),
    CONSTRAINT tag_color_check CHECK (color LIKE'%[0-9A-Fa-f]%') NOT VALID
)

---- create above / drop below ----

DROP TABLE IF EXISTS public.tag;
