CREATE TABLE public.users
(
    id serial NOT NULL,
    login character varying(64) NOT NULL,
    password character varying(128) NOT NULL,
    user_role character varying(32) NOT NULL,
    PRIMARY KEY (id)
);

ALTER TABLE public.users
    OWNER to postgres;