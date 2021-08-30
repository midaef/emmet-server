create table public.values
(
    id serial not null
        constraint values_pk
            primary key,
    value text not null,
    key text not null,
    allowed_roles int array not null,
    created_by int not null,
    config int not null
);

ALTER TABLE public.roles
    OWNER to postgres;
