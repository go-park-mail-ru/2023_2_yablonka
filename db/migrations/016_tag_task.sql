CREATE TABLE IF NOT EXISTS public.Tag_Task
(
    id_tag serial NOT NULL,
    id_task serial NOT NULL,
    PRIMARY KEY (id_tag, id_task)
        INCLUDE(id_tag, id_task),
    UNIQUE (id_tag, id_task)
        INCLUDE(id_tag, id_task)
);

ALTER TABLE IF EXISTS public.Tag_Task
    ADD FOREIGN KEY (id_tag)
    REFERENCES public.Tag (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;


ALTER TABLE IF EXISTS public.Tag_Task
    ADD FOREIGN KEY (id_task)
    REFERENCES public.Task (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;


---- create above / drop below ----

DROP TABLE IF EXISTS public.Tag_Task;
