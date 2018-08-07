-- ----------------------------
-- Table structure for alliances
-- ----------------------------
DROP TABLE IF EXISTS alliances;

CREATE TABLE alliances (
  name varchar(20) COLLATE "default" NOT NULL,
  created_at timestamp(6) DEFAULT now() NOT NULL,
  updated_at timestamp(6) DEFAULT now() NOT NULL
)
WITH (OIDS = FALSE);

-- ----------------------------
-- Table structure for cities
-- ----------------------------
DROP TABLE IF EXISTS cities;

CREATE TABLE cities (
  x int2 CHECK (x < 600
    AND x > - 1) NOT NULL,
  y int2 CHECK (y < 600
    AND y > - 1) NOT NULL,
  continent_x int2 NOT NULL,
  continent_y int2 NOT NULL,
  TYPE int2 DEFAULT 1 NOT NULL,
  created_at timestamp(6) DEFAULT now() NOT NULL,
  updated_at timestamp(6) DEFAULT now() NOT NULL,
  user_name varchar(20) COLLATE "default",
  name varchar(20) COLLATE "default" DEFAULT 'New City' ::character varying NOT NULL,
  points int2 DEFAULT 3 NOT NULL,
  wood_production int DEFAULT 500 NOT NULL,
  stone_production int DEFAULT 0 NOT NULL,
  iron_production int DEFAULT 0 NOT NULL,
  food_production int DEFAULT 0 NOT NULL,
  gold_production int DEFAULT 0 NOT NULL,
  wood_stored int DEFAULT 1000 CHECK (wood_stored <= wood_limit
    AND wood_stored >= 0) NOT NULL,
  stone_stored int DEFAULT 1000 CHECK (stone_stored <= stone_limit
    AND stone_stored >= 0) NOT NULL,
  iron_stored int DEFAULT 500 CHECK (iron_stored <= iron_limit
    AND iron_stored >= 0) NOT NULL,
  food_stored int DEFAULT 500 CHECK (food_stored <= food_limit
    AND food_stored >= 0) NOT NULL,
  wood_limit int DEFAULT 5000 NOT NULL,
  stone_limit int DEFAULT 5000 NOT NULL,
  iron_limit int DEFAULT 5000 NOT NULL,
  food_limit int DEFAULT 5000 NOT NULL
)
WITH (OIDS = FALSE);

-- ----------------------------
-- Table structure for constructions
-- ----------------------------
DROP TABLE IF EXISTS constructions;

-- auto-generated definition
CREATE TABLE constructions (
  x int2 CHECK (x <= 20
    AND x >= 0) NOT NULL,
  y int2 CHECK (y <= 20
    AND y >= 0) NOT NULL,
  city_x int2 CHECK (city_x <= 600
    AND city_x >= 0) NOT NULL,
  city_y int2 CHECK (city_y <= 600
    AND city_y >= 0) NOT NULL,
  created_at timestamp(6) DEFAULT now() NOT NULL,
  updated_at timestamp(6) DEFAULT now() NOT NULL,
  level int2 DEFAULT 0 CHECK (level >= 1
    AND level <= 10) NOT NULL,
  TYPE int2 NOT NULL,
  production int2 DEFAULT 0 NOT NULL,
  modifier int2 DEFAULT 1 NOT NULL,
  need_refresh boolean DEFAULT TRUE NOT NULL
)
WITH (OIDS = FALSE);

-- ----------------------------
-- Table structure for continents
-- ----------------------------
DROP TABLE IF EXISTS continents;

CREATE TABLE continents (
  x int2 NOT NULL,
  y int2 NOT NULL,
  created_at timestamp(6) DEFAULT now() NOT NULL,
  updated_at timestamp(6) DEFAULT now() NOT NULL,
  is_active bool DEFAULT FALSE NOT NULL,
  size int2 DEFAULT 100 NOT NULL,
  number_of_cities int2 DEFAULT 0 NOT NULL,
  cities_limit int2 DEFAULT 1000 NOT NULL
)
WITH (OIDS = FALSE);

-- ----------------------------
-- Table structure for dungeons
-- ----------------------------
DROP TABLE IF EXISTS dungeons;

CREATE TABLE dungeons (
  x int2 CHECK (x <= 600
    AND x >= 0) NOT NULL,
  y int2 CHECK (y <= 600
    AND y >= 0) NOT NULL,
  continent_x int2 NOT NULL,
  continent_y int2 NOT NULL,
  TYPE int2 NOT NULL,
  level int2 CHECK (level <= 10
    AND level >= 0) DEFAULT 1 NOT NULL,
  progress int2 CHECK (progress <= 100
    AND progress >= 0) DEFAULT 0 NOT NULL
)
WITH (OIDS = FALSE);

-- ----------------------------
-- Table structure for military_actions
-- ----------------------------
DROP TABLE IF EXISTS military_actions;

CREATE TABLE military_actions (
  id int NOT NULL,
  origin_id int,
  target_id int,
  arrival timestamp(6) NOT NULL,
  troops json NOT NULL
)
WITH (OIDS = FALSE);

-- ----------------------------
-- Table structure for tiles
-- ----------------------------
DROP TABLE IF EXISTS tiles;

CREATE TABLE tiles (
  x int2 CHECK (x < 600
    AND x > - 1) NOT NULL,
  y int2 CHECK (y < 600
    AND y > - 1) NOT NULL,
  continent_x int2 NOT NULL,
  continent_y int2 NOT NULL,
  occupied_by varchar(20) COLLATE "default" DEFAULT 'land' ::character varying NOT NULL
)
WITH (OIDS = FALSE);

-- ----------------------------
-- Table structure for upgrades
-- ----------------------------
DROP TABLE IF EXISTS upgrades;

CREATE TABLE upgrades (
  x int2 CHECK (x <= 20
    AND x >= 0) NOT NULL,
  y int2 CHECK (y <= 20
    AND y >= 0) NOT NULL,
  city_x int2 CHECK (city_x <= 600
    AND city_x >= 0) NOT NULL,
  city_y int2 CHECK (city_y <= 600
    AND city_y >= 0) NOT NULL,
  created_at timestamp(6) DEFAULT now() NOT NULL,
  INDEX int2 DEFAULT 0 NOT NULL,
  duration int DEFAULT 10 NOT NULL,
  START timestamp(0) DEFAULT now() NOT NULL
)
WITH (OIDS = FALSE);

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS users;

CREATE TABLE users (
  created_at timestamp(6) DEFAULT now() NOT NULL,
  updated_at timestamp(6) DEFAULT now() NOT NULL,
  name varchar(20) COLLATE "default" NOT NULL,
  email varchar(40) COLLATE "default" NOT NULL,
  PASSWORD varchar(100) COLLATE "default" NOT NULL,
  gold int DEFAULT 1000 NOT NULL,
  diamonds int DEFAULT 0 NOT NULL,
  darkwood int DEFAULT 0 NOT NULL,
  runestone int DEFAULT 0 NOT NULL,
  veritium int DEFAULT 0 NOT NULL,
  trueseed int DEFAULT 0 NOT NULL,
  rank int2 DEFAULT 0 NOT NULL,
  alliance_name varchar(20) COLLATE "default",
  alliance_rank varchar(15) COLLATE "default"
)
WITH (OIDS = FALSE);

-- ----------------------------
-- Alter Sequences Owned By 
-- ----------------------------
-- ----------------------------
-- Uniques structure for table alliances
-- ----------------------------
ALTER TABLE alliances
  ADD UNIQUE (name);

-- ----------------------------
-- Indexes structure for table cities
-- ----------------------------
CREATE UNIQUE INDEX city_coord_index ON cities
USING btree (x, y, continent_x, continent_y);

ALTER TABLE cities CLUSTER ON city_coord_index;

CREATE INDEX city_username_index ON cities
USING btree (user_name);

-- ----------------------------
-- Uniques structure for table cities
-- ----------------------------
ALTER TABLE cities
  ADD UNIQUE (x,
    y);

-- ----------------------------
-- Primary Key structure for table cities
-- ----------------------------
ALTER TABLE cities
  ADD PRIMARY KEY (x, y, continent_x, continent_y);

-- ----------------------------
-- Indexes structure for table constructions
-- ----------------------------
CREATE UNIQUE INDEX city_index ON constructions
USING btree (x, y, city_x, city_y);

ALTER TABLE constructions CLUSTER ON city_index;

-- ----------------------------
-- Uniques structure for table constructions
-- ----------------------------
ALTER TABLE constructions
  ADD UNIQUE (x,
    y,
    city_x,
    city_y);

-- ----------------------------
-- Primary Key structure for table constructions
-- ----------------------------
ALTER TABLE constructions
  ADD PRIMARY KEY (city_x, x, y, city_y);

-- ----------------------------
-- Indexes structure for table continents
-- ----------------------------
CREATE UNIQUE INDEX xy_index ON continents
USING btree (x, y);

ALTER TABLE continents CLUSTER ON xy_index;

-- ----------------------------
-- Uniques structure for table continents
-- ----------------------------
ALTER TABLE continents
  ADD UNIQUE (x,
    y);

-- ----------------------------
-- Primary Key structure for table continents
-- ----------------------------
ALTER TABLE continents
  ADD PRIMARY KEY (x, y);

-- ----------------------------
-- Indexes structure for table dungeons
-- ----------------------------
CREATE UNIQUE INDEX coord_dungeon_index ON dungeons
USING btree (continent_x, continent_y, x, y);

ALTER TABLE dungeons CLUSTER ON coord_dungeon_index;

-- ----------------------------
-- Uniques structure for table dungeons
-- ----------------------------
ALTER TABLE dungeons
  ADD UNIQUE (x,
    y);

-- ----------------------------
-- Primary Key structure for table dungeons
-- ----------------------------
ALTER TABLE dungeons
  ADD PRIMARY KEY (continent_x, continent_y, x, y);

-- ----------------------------
-- Indexes structure for table military_actions
-- ----------------------------
CREATE INDEX origin_index ON military_actions
USING btree (origin_id);

CREATE INDEX target_index ON military_actions
USING btree (target_id);

CREATE INDEX arrival_index ON military_actions
USING btree (arrival);

-- ----------------------------
-- Primary Key structure for table military_actions
-- ----------------------------
ALTER TABLE military_actions
  ADD PRIMARY KEY (id);

-- ----------------------------
-- Indexes structure for table tiles
-- ----------------------------
CREATE UNIQUE INDEX coord_tile_index ON tiles
USING btree (x, y, continent_x, continent_y);

ALTER TABLE tiles CLUSTER ON coord_tile_index;

-- ----------------------------
-- Uniques structure for table tiles
-- ----------------------------
ALTER TABLE tiles
  ADD UNIQUE (x,
    y);

-- ----------------------------
-- Primary Key structure for table tiles
-- ----------------------------
ALTER TABLE tiles
  ADD PRIMARY KEY (x, y, continent_x, continent_y);

-- ----------------------------
-- Indexes structure for table upgrades
-- ----------------------------
CREATE INDEX order_index ON upgrades
USING btree (INDEX);

CREATE UNIQUE INDEX tile_upgrade_index ON upgrades
USING btree (x, y, city_x, city_y, INDEX);

ALTER TABLE upgrades CLUSTER ON tile_upgrade_index;

-- ----------------------------
-- Uniques structure for table upgrades
-- ----------------------------
ALTER TABLE upgrades
  ADD UNIQUE (x,
    y,
    city_x,
    city_y,
    INDEX);

-- ----------------------------
-- Primary Key structure for table upgrades
-- ----------------------------
ALTER TABLE upgrades
  ADD PRIMARY KEY (x, y, city_x, city_y);

-- ----------------------------
-- Indexes structure for table users
-- ----------------------------
CREATE UNIQUE INDEX email_index ON users
USING btree (email);

CREATE UNIQUE INDEX user_name_index ON users
USING btree (name);

-- ----------------------------
-- Uniques structure for table users
-- ----------------------------
ALTER TABLE users
  ADD UNIQUE (name);

ALTER TABLE users
  ADD UNIQUE (email);

-- ----------------------------
-- Primary Key structure for table users
-- ----------------------------
ALTER TABLE users
  ADD PRIMARY KEY (email, name);

-- ----------------------------
-- Foreign Key structure for table cities
-- ----------------------------
ALTER TABLE cities
  ADD FOREIGN KEY (x,
    y,
    continent_x,
    continent_y)
  REFERENCES tiles (x,
    y,
    continent_x,
    continent_y)
  ON DELETE CASCADE ON
  UPDATE
    CASCADE;

ALTER TABLE cities
  ADD FOREIGN KEY (user_name)
  REFERENCES users (name)
  ON DELETE SET NULL ON
  UPDATE
    CASCADE;

-- ----------------------------
-- Foreign Key structure for table constructions
-- ----------------------------
ALTER TABLE constructions
  ADD FOREIGN KEY (city_x,
    city_y)
  REFERENCES cities (x,
    y)
  ON DELETE CASCADE ON
  UPDATE
    CASCADE;

-- ----------------------------
-- Foreign Key structure for table dungeons
-- ----------------------------
ALTER TABLE dungeons
  ADD FOREIGN KEY (x,
    y,
    continent_x,
    continent_y)
  REFERENCES tiles (x,
    y,
    continent_x,
    continent_y)
  ON DELETE CASCADE ON
  UPDATE
    CASCADE;

-- ----------------------------
-- Foreign Key structure for table tiles
-- ----------------------------
ALTER TABLE tiles
  ADD FOREIGN KEY (continent_x,
    continent_y)
  REFERENCES continents (x,
    y)
  ON DELETE CASCADE ON
  UPDATE
    CASCADE;

-- ----------------------------
-- Foreign Key structure for table upgrades
-- ----------------------------
ALTER TABLE upgrades
  ADD FOREIGN KEY (x,
    y,
    city_x,
    city_y)
  REFERENCES constructions (x,
    y,
    city_x,
    city_y)
  ON DELETE CASCADE ON
  UPDATE
    CASCADE;

-- ----------------------------
-- Foreign Key structure for table users
-- ----------------------------
ALTER TABLE users
  ADD FOREIGN KEY (alliance_name)
  REFERENCES alliances (name)
  ON DELETE SET NULL ON
  UPDATE
    CASCADE;

