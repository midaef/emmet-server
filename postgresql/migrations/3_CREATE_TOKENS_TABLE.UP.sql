CREATE TABLE public.tokens
(
    id serial NOT NULL,
    token character varying(512) NOT NULL,
    exp time without time zone NOT NULL,
    user_login character varying(64) NOT NULL,
    user_role character varying(32) NOT NULL,
    PRIMARY KEY (id)
);

ALTER TABLE public.tokens
    OWNER to postgres;