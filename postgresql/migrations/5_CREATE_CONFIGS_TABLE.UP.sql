create table public.configs
(
    id serial not null
        constraint configs_pk
            primary key,
    config_name varchar(32) not null,
    created_by int not null,
    allowed_roles int array not null,
    internal_configs int array
);

ALTER TABLE public.roles
    OWNER to postgres;