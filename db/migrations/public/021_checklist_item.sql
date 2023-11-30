CREATE TABLE IF NOT EXISTS public.checklist_item
(
    id serial NOT NULL,
    id_checklist serial NOT NULL,
    name text NOT NULL,
    done boolean NOT NULL DEFAULT false,
    list_position smallint NOT NULL,
    CONSTRAINT checklist_item_pkey PRIMARY KEY (id),
    CONSTRAINT checklist_item_id_checklist_fkey FOREIGN KEY (id_checklist)
        REFERENCES public.checklist (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE CASCADE
        NOT VALID,
    CONSTRAINT checklist_item_list_position_check CHECK (list_position >= 0) NOT VALID
)

---- create above / drop below ----

DROP TABLE IF EXISTS public.checklist_item;
