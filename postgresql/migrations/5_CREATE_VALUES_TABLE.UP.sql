CREATE TABLE public."values"
(
    id serial NOT NULL,
    created_by character varying(32) NOT NULL,
    key text NOT NULL,
    value text NOT NULL,
    roles text NOT NULL,
    PRIMARY KEY (id)
);

ALTER TABLE public."values"
    OWNER to postgres;