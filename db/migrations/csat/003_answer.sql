CREATE TABLE IF NOT EXISTS csat_db.answer
(
    id integer NOT NULL DEFAULT nextval('csat_db.response_id_seq'::regclass),
    id_user integer NOT NULL DEFAULT nextval('csat_db.response_id_user_seq'::regclass),
    score smallint NOT NULL DEFAULT 0,
    id_question integer NOT NULL DEFAULT nextval('csat_db.response_id_question_seq'::regclass),
    CONSTRAINT response_pkey PRIMARY KEY (id),
    CONSTRAINT answer_id_question_fkey FOREIGN KEY (id_question)
        REFERENCES csat_db.question (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
        NOT VALID,
    CONSTRAINT response_id_user_fkey FOREIGN KEY (id_user)
        REFERENCES public."user" (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION,
    CONSTRAINT response_score_left_limit_check CHECK (1 <= score),
    CONSTRAINT response_score_right_limit_check CHECK (score <= 5) NOT VALID
)

---- create above / drop below ----

DROP TABLE IF EXISTS csat_db.answer;