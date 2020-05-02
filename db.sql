create database openlou
with owner postgres;

create table alliances
(
  name       varchar(20)                not null,
  created_at timestamp(6) default now() not null,
  updated_at timestamp(6) default now() not null,
  id         serial                     not null
    constraint alliances_pk
    primary key
);

alter table alliances
  owner to postgres;

create unique index alliances_id_uindex
  on alliances (id);

create table cities
(
  x                integer                                             not null,
  y                integer                                             not null,
  type             integer default 1                                   not null,
  created_at       timestamp(6) default now()                          not null,
  updated_at       timestamp(6) default now()                          not null,
  user_name        varchar(20),
  city_name        varchar(20) default 'New City' :: character varying not null,
  points           integer default 3                                   not null,
  wood_production  integer default 500                                 not null,
  stone_production integer default 0                                   not null,
  iron_production  integer default 0                                   not null,
  food_production  integer default 0                                   not null,
  gold_production  integer default 0                                   not null,
  wood_stored      integer default 1000                                not null,
  stone_stored     integer default 1000                                not null,
  iron_stored      integer default 500                                 not null,
  food_stored      integer default 500                                 not null,
  wood_limit       integer default 5000                                not null,
  stone_limit      integer default 5000                                not null,
  iron_limit       integer default 5000                                not null,
  food_limit       integer default 5000                                not null,
  queue_time       timestamp default now()                                not null,
  construction_speed integer default '100'                             not null,
  constraint city_pkey
  primary key (x, y)
);

alter table cities
  owner to postgres;

create table constructions
(
  x            integer                    not null,
  y            integer                    not null,
  city_x       integer                    not null,
  city_y       integer                    not null,
  created_at   timestamp(6) default now() not null,
  updated_at   timestamp(6) default now() not null,
  level        integer default 0          not null,
  type         integer                    not null,
  production   integer default 0          not null,
  modifier     integer default 1          not null,
  need_refresh boolean default true       not null,
  constraint construction_pkey
  primary key (city_x, city_y, x, y)
);

alter table constructions
  owner to postgres;

create table dungeons
(
  x        integer           not null,
  y        integer           not null,
  type     integer           not null,
  level    integer default 1 not null,
  progress integer default 0 not null,
  constraint dungeon_pkey
  primary key (x, y)
);

alter table dungeons
  owner to postgres;

create table military_actions
(
  id        integer      not null
    constraint military_action_pkey
    primary key,
  origin_id integer,
  target_id integer,
  arrival   timestamp(6) not null,
  troops    json         not null
);

alter table military_actions
  owner to postgres;

create table queue
(
  construction_x integer                    not null,
  construction_y integer                    not null,
  city_x         integer                    not null,
  city_y         integer                    not null,
  created_at     timestamp(6) default now() not null,
  completion     timestamp(0) default now() not null,
  action         integer                    not null,
  constraint upgrade_pkey
  primary key (construction_x, construction_y, city_x, city_y)
);

alter table queue
  owner to postgres;

create table users
(
    created_at    timestamp(6) default now() not null,
    updated_at    timestamp(6) default now() not null,
    name          varchar(20)                not null
        constraint users_pk
            primary key,
    email         varchar(40)                not null
        constraint users_email_key
            unique,
    password      varchar(100)               not null,
    gold          integer      default 1000  not null,
    diamonds      integer      default 0     not null,
    darkwood      integer      default 0     not null,
    runestone     integer      default 0     not null,
    veritium      integer      default 0     not null,
    trueseed      integer      default 0     not null,
    rank          integer      default 0     not null,
    alliance_name varchar(20),
    alliance_rank varchar(15)
);

alter table users
    owner to postgres;



