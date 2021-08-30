create table public.roles
(
    id serial not null
        constraint roles_pk
            primary key,
    role_name varchar(32) not null,
    created_by int not null,
    create_user bool not null,
    delete_user bool not null,
    update_user bool not null,
    create_config bool not null,
    delete_config bool not null,
    update_config bool not null,
    create_role bool not null,
    delete_role bool not null,
    update_role bool not null,
    create_value bool not null,
    delete_value bool not null,
    update_value bool not null,
    allowed_users int array
);

ALTER TABLE public.roles
    OWNER to postgres;