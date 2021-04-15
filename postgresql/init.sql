CREATE TABLE public.users
(
    id serial NOT NULL,
    login character varying(64) NOT NULL,
    password character varying(128) NOT NULL,
    user_role character varying(32) NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE public.tokens
(
    id serial NOT NULL,
    token character varying(512) NOT NULL,
    exp time without time zone NOT NULL,
    user_login character varying(64) NOT NULL,
    user_role character varying(32) NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE public.roles
(
    id serial NOT NULL,
    created_by character varying(32) NOT NULL,
    create_user boolean NOT NULL,
    create_role boolean NOT NULL,
    create_value boolean NOT NULL,
    delete_user boolean NOT NULL,
    delete_value boolean NOT NULL,
    delete_role boolean NOT NULL,
    user_role character varying(32) NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE public."values"
(
    id serial NOT NULL,
    created_by character varying(32) NOT NULL,
    key text NOT NULL,
    value text NOT NULL,
    roles text NOT NULL,
    PRIMARY KEY (id)
);

INSERT INTO public.roles
(
    created_by,
    create_user,
    create_role,
    create_value,
    user_role,
    delete_user,
    delete_role,
    delete_value
) VALUES (
    'root',
    true,
    true,
    true,
    'root',
    true,
    true,
    true
);

INSERT INTO public.users
(
    login,
    password,
    user_role
)
VALUES
(
    'root',
    '5f4dcc3b5aa765d61d8327deb882cf9973616c74',
    'root'
);

COMMIT;
