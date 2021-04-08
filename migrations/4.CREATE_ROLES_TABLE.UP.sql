CREATE TABLE public.roles
(
    id serial NOT NULL,
    created_by character varying(32) NOT NULL,
    create_user boolean NOT NULL,
    create_role boolean NOT NULL,
    create_value boolean NOT NULL,
    user_role character varying(32) NOT NULL,
    PRIMARY KEY (id)
);

ALTER TABLE public.roles
    OWNER to postgres;