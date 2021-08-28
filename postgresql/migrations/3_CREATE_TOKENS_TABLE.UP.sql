CREATE TABLE public.tokens
(
    id            SERIAL PRIMARY KEY,
    user_id       INTEGER,
    access_token  VARCHAR(600),
    refresh_token VARCHAR(64),
    exp           TIMESTAMP NOT NULL
);

ALTER TABLE public.tokens
    OWNER to postgres;
