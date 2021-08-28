CREATE TABLE public.users
(
    id serial NOT NULL,
    login character varying(32) NOT NULL,
    password character varying(64) NOT NULL,
    user_role character varying(32) NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE public.tokens
(
    id            SERIAL PRIMARY KEY,
    user_id       INTEGER,
    access_token  VARCHAR(600),
    refresh_token VARCHAR(64),
    exp           TIMESTAMP NOT NULL
);
