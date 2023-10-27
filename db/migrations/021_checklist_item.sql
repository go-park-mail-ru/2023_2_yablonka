CREATE TABLE IF NOT EXISTS public.Checklist_item
(
    id serial NOT NULL,
    id_checklist serial NOT NULL,
    name text NOT NULL,
    done boolean NOT NULL DEFAULT false,
    list_position smallint NOT NULL,
    PRIMARY KEY (id),
    UNIQUE (id)
        INCLUDE(id)
);

ALTER TABLE IF EXISTS public.Checklist_item
    ADD FOREIGN KEY (id_checklist)
    REFERENCES public.Checklist (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;


---- create above / drop below ----

DROP TABLE IF EXISTS public.Checklist_item;
