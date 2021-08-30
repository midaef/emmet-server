CREATE TABLE public.users
(
    id serial NOT NULL,
    login character varying(32) NOT NULL,
    password character varying(64) NOT NULL,
    user_role character varying(32) NOT NULL,
    created_by int not null,
    PRIMARY KEY (id)
);

ALTER TABLE public.users
    OWNER to postgres;