CREATE TABLE public.cities (
  id      INT4        NOT NULL DEFAULT nextval('cities_id_seq' :: REGCLASS),
  "name"  VARCHAR(15) NOT NULL,
  x       INT4        NOT NULL,
  y       INT4        NOT NULL,
  points  INT4        NOT NULL DEFAULT 3,
  user_id INT4        NULL,
  CONSTRAINT cities_pkey PRIMARY KEY (id),
  CONSTRAINT user_fk FOREIGN KEY (user_id) REFERENCES public.users (id) ON DELETE SET NULL
)
WITH (
OIDS = FALSE
);
CREATE UNIQUE INDEX cities_id_uindex
  ON public.cities (id int4_ops);
CREATE INDEX cities_user_index
  ON public.cities (user_id int4_ops);

CREATE TABLE public.constructions (
  id           INT4 NOT NULL DEFAULT nextval('constructions_id_seq' :: REGCLASS),
  city_id      INT4 NOT NULL,
  "level"      INT4 NOT NULL DEFAULT 1,
  x            INT4 NOT NULL,
  y            INT4 NOT NULL,
  build_type   INT4 NOT NULL,
  need_refresh BOOL NOT NULL DEFAULT TRUE,
  CONSTRAINT constructions_pkey PRIMARY KEY (id),
  CONSTRAINT city_fk FOREIGN KEY (city_id) REFERENCES public.cities (id) ON DELETE CASCADE
)
WITH (
OIDS = FALSE
);
CREATE INDEX constructions_city_id_index
  ON public.constructions (city_id int4_ops);
CREATE UNIQUE INDEX constructions_id_uindex
  ON public.constructions (id int4_ops);

CREATE TABLE public.military (
  id        INT4 NOT NULL DEFAULT nextval('military_id_seq' :: REGCLASS),
  city_from INT4 NOT NULL,
  city_to   INT4 NULL,
  "type"    INT4 NOT NULL,
  troops    JSON NOT NULL,
  CONSTRAINT military_pkey PRIMARY KEY (id),
  CONSTRAINT city_from_fk FOREIGN KEY (city_from) REFERENCES public.cities (id) ON DELETE CASCADE,
  CONSTRAINT city_to_fk FOREIGN KEY (city_to) REFERENCES public.cities (id) ON DELETE SET NULL
)
WITH (
OIDS = FALSE
);
CREATE UNIQUE INDEX military_id_uindex
  ON public.military (id int4_ops);

CREATE TABLE public.reports (
  id      INT4    NOT NULL DEFAULT nextval('reports_id_seq' :: REGCLASS),
  readers INT4 [] NOT NULL,
  city_id INT4    NOT NULL,
  info    JSON    NOT NULL,
  CONSTRAINT reports_pkey PRIMARY KEY (id),
  CONSTRAINT city_fk FOREIGN KEY (city_id) REFERENCES public.cities (id) ON DELETE CASCADE
)
WITH (
OIDS = FALSE
);
CREATE UNIQUE INDEX reports_id_uindex
  ON public.reports (id int4_ops);

CREATE TABLE public.upgrades (
  id              INT4      NOT NULL DEFAULT nextval('upgrades_id_seq' :: REGCLASS),
  construction_id INT4      NOT NULL,
  city_id         INT4      NOT NULL,
  completion      TIMESTAMP NOT NULL,
  CONSTRAINT upgrades_pkey PRIMARY KEY (id),
  CONSTRAINT city_fk FOREIGN KEY (city_id) REFERENCES public.cities (id),
  CONSTRAINT construction_fk FOREIGN KEY (construction_id) REFERENCES public.constructions (id) ON DELETE CASCADE
)
WITH (
OIDS = FALSE
);
CREATE UNIQUE INDEX upgrades_id_uindex
  ON public.upgrades (id int4_ops);

CREATE TABLE public.users (
  id       INT4         NOT NULL DEFAULT nextval('users_id_seq' :: REGCLASS),
  login    VARCHAR(15)  NOT NULL,
  password VARCHAR(100) NOT NULL,
  email    VARCHAR(30)  NOT NULL,
  CONSTRAINT users_id_pk PRIMARY KEY (id),
  CONSTRAINT users_mail_un UNIQUE (email),
  CONSTRAINT users_un UNIQUE (login)
)
WITH (
OIDS = FALSE
);
CREATE UNIQUE INDEX users_email_idx
  ON public.users (email text_ops);
CREATE UNIQUE INDEX users_id_uindex
  ON public.users (id int4_ops);
CREATE UNIQUE INDEX users_login_uindex
  ON public.users (login text_ops);

CREATE TABLE public.sessions (
  "key"   VARCHAR(64) NOT NULL,
  user_id INT4        NOT NULL,
  ip      INET        NOT NULL,
  CONSTRAINT sessions_pk PRIMARY KEY ("key")
)
WITH (
OIDS = FALSE
);
CREATE UNIQUE INDEX sessions_key_idx
  ON public.sessions ("key" text_ops);
