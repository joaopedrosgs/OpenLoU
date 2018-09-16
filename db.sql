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
  x                integer                                             not null
    constraint city_x_check
    check ((x < 600) AND (x > '-1' :: integer)),
  y                integer                                             not null
    constraint city_y_check
    check ((y < 600) AND (y > '-1' :: integer)),
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
  constraint city_pkey
  primary key (x, y),
  constraint city_check
  check ((wood_stored <= wood_limit) AND (wood_stored >= 0)),
  constraint city_check1
  check ((stone_stored <= stone_limit) AND (stone_stored >= 0)),
  constraint city_check2
  check ((iron_stored <= iron_limit) AND (iron_stored >= 0)),
  constraint city_check3
  check ((food_stored <= food_limit) AND (food_stored >= 0))
);

alter table cities
  owner to postgres;

create table constructions
(
  x            integer                    not null
    constraint construction_x_check
    check ((x <= 20) AND (x >= 0)),
  y            integer                    not null
    constraint construction_y_check
    check ((y <= 20) AND (y >= 0)),
  city_x       integer                    not null
    constraint construction_city_x_check
    check ((city_x <= 600) AND (city_x >= 0)),
  city_y       integer                    not null
    constraint construction_city_y_check
    check ((city_y <= 600) AND (city_y >= 0)),
  created_at   timestamp(6) default now() not null,
  updated_at   timestamp(6) default now() not null,
  level        integer default 0          not null
    constraint construction_level_check
    check ((level >= 1) AND (level <= 10)),
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
  x        integer           not null
    constraint dungeon_x_check
    check ((x <= 600) AND (x >= 0)),
  y        integer           not null
    constraint dungeon_y_check
    check ((y <= 600) AND (y >= 0)),
  type     integer           not null,
  level    integer default 1 not null
    constraint dungeon_level_check
    check ((level <= 10) AND (level >= 0)),
  progress integer default 0 not null
    constraint dungeon_progress_check
    check ((progress <= 100) AND (progress >= 0)),
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

create table upgrades
(
  construction_x integer                    not null
    constraint upgrade_construction_x_check
    check ((construction_x <= 20) AND (construction_x >= 0)),
  construction_y integer                    not null
    constraint upgrade_construction_y_check
    check ((construction_y <= 20) AND (construction_y >= 0)),
  city_x         integer                    not null
    constraint upgrade_city_x_check
    check ((city_x <= 600) AND (city_x >= 0)),
  city_y         integer                    not null
    constraint upgrade_city_y_check
    check ((city_y <= 600) AND (city_y >= 0)),
  created_at     timestamp(6) default now() not null,
  index_at_queue integer default 0          not null,
  duration       integer default 10         not null,
  start          timestamp(0) default now() not null,
  constraint upgrade_pkey
  primary key (construction_x, construction_y, city_x, city_y)
);

alter table upgrades
  owner to postgres;

create table users
(
  created_at    timestamp(6) default now() not null,
  updated_at    timestamp(6) default now() not null,
  name          varchar(20)                not null,
  email         varchar(40)                not null
    constraint user_pkey
    primary key,
  password      varchar(100)               not null,
  gold          integer default 1000       not null,
  diamonds      integer default 0          not null,
  darkwood      integer default 0          not null,
  runestone     integer default 0          not null,
  veritium      integer default 0          not null,
  trueseed      integer default 0          not null,
  rank          integer default 0          not null,
  alliance_name varchar(20),
  alliance_rank varchar(15)
);

alter table users
  owner to postgres;

