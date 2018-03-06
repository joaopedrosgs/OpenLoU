/*
Navicat PGSQL Data Transfer

Source Server         : local
Source Server Version : 100100
Source Host           : localhost:5432
Source Database       : openlou
Source Schema         : public

Target Server Type    : PGSQL
Target Server Version : 100100
File Encoding         : 65001

Date: 2018-03-06 07:54:13
*/


-- ----------------------------
-- Table structure for alliances
-- ----------------------------
DROP TABLE IF EXISTS "public"."alliances";
CREATE TABLE "public"."alliances" (
"name" varchar(20) COLLATE "default" NOT NULL,
"created_at" timestamp(6) DEFAULT now() NOT NULL,
"updated_at" timestamp(6) DEFAULT now() NOT NULL
)
WITH (OIDS=FALSE)

;

-- ----------------------------
-- Table structure for cities
-- ----------------------------
DROP TABLE IF EXISTS "public"."cities";
CREATE TABLE "public"."cities" (
"x" int2 NOT NULL,
"y" int2 NOT NULL,
"continent_x" int2 NOT NULL,
"continent_y" int2 NOT NULL,
"type" int2 DEFAULT 1 NOT NULL,
"created_at" timestamp(6) DEFAULT now() NOT NULL,
"updated_at" timestamp(6) DEFAULT now() NOT NULL,
"user_name" varchar(20) COLLATE "default",
"name" varchar(20) COLLATE "default" DEFAULT 'New City *'::character varying NOT NULL,
"points" int2 DEFAULT 3 NOT NULL,
"wood_production" int4 DEFAULT 500 NOT NULL,
"stone_production" int4 DEFAULT 0 NOT NULL,
"iron_production" int4 DEFAULT 0 NOT NULL,
"food_production" int4 DEFAULT 0 NOT NULL,
"gold_production" int4 DEFAULT 0 NOT NULL,
"wood_stored" int4 DEFAULT 1000 NOT NULL,
"stone_stored" int4 DEFAULT 1000 NOT NULL,
"iron_stored" int4 DEFAULT 500 NOT NULL,
"food_stored" int4 DEFAULT 500 NOT NULL,
"wood_limit" int4 DEFAULT 5000 NOT NULL,
"stone_limit" int4 DEFAULT 5000 NOT NULL,
"iron_limit" int4 DEFAULT 5000 NOT NULL,
"food_limit" int4 DEFAULT 5000 NOT NULL
)
WITH (OIDS=FALSE)

;

-- ----------------------------
-- Table structure for constructions
-- ----------------------------
DROP TABLE IF EXISTS "public"."constructions";
CREATE TABLE "public"."constructions" (
"x" int2 NOT NULL,
"y" int2 NOT NULL,
"city_x" int2 NOT NULL,
"city_y" int2 NOT NULL,
"created_at" timestamp(6) DEFAULT now() NOT NULL,
"updated_at" timestamp(6) DEFAULT now() NOT NULL,
"level" int2 DEFAULT 0 NOT NULL,
"type" int2 NOT NULL,
"production" int2 DEFAULT 0 NOT NULL,
"modifier" int2 DEFAULT 1 NOT NULL,
"need_refresh" bool DEFAULT true NOT NULL
)
WITH (OIDS=FALSE)

;

-- ----------------------------
-- Table structure for continents
-- ----------------------------
DROP TABLE IF EXISTS "public"."continents";
CREATE TABLE "public"."continents" (
"x" int2 NOT NULL,
"y" int2 NOT NULL,
"created_at" timestamp(6) DEFAULT now() NOT NULL,
"updated_at" timestamp(6) DEFAULT now() NOT NULL,
"is_active" bool DEFAULT false NOT NULL,
"size" int2 DEFAULT 100 NOT NULL,
"number_of_cities" int2 DEFAULT 0 NOT NULL,
"cities_limit" int2 DEFAULT 1000 NOT NULL
)
WITH (OIDS=FALSE)

;

-- ----------------------------
-- Table structure for dungeons
-- ----------------------------
DROP TABLE IF EXISTS "public"."dungeons";
CREATE TABLE "public"."dungeons" (
"x" int2 NOT NULL,
"y" int2 NOT NULL,
"continent_x" int2 NOT NULL,
"continent_y" int2 NOT NULL,
"type" int2 NOT NULL,
"level" int2 DEFAULT 1 NOT NULL,
"progress" int2 DEFAULT 0 NOT NULL
)
WITH (OIDS=FALSE)

;

-- ----------------------------
-- Table structure for military_actions
-- ----------------------------
DROP TABLE IF EXISTS "public"."military_actions";
CREATE TABLE "public"."military_actions" (
"id" int4 NOT NULL,
"origin_id" int4,
"target_id" int4,
"arrival" timestamp(6) NOT NULL,
"troops" json NOT NULL
)
WITH (OIDS=FALSE)

;

-- ----------------------------
-- Table structure for tiles
-- ----------------------------
DROP TABLE IF EXISTS "public"."tiles";
CREATE TABLE "public"."tiles" (
"x" int2 NOT NULL,
"y" int2 NOT NULL,
"continent_x" int2 NOT NULL,
"continent_y" int2 NOT NULL,
"occupied_by" varchar(20) COLLATE "default" DEFAULT 'land'::character varying NOT NULL
)
WITH (OIDS=FALSE)

;

-- ----------------------------
-- Table structure for upgrades
-- ----------------------------
DROP TABLE IF EXISTS "public"."upgrades";
CREATE TABLE "public"."upgrades" (
"x" int2 NOT NULL,
"y" int2 NOT NULL,
"city_x" int4 NOT NULL,
"city_y" int2 NOT NULL,
"created_at" timestamp(6) DEFAULT now() NOT NULL,
"updated_at" date DEFAULT now() NOT NULL,
"index" int2 DEFAULT 1 NOT NULL,
"duration" int4 DEFAULT 10 NOT NULL,
"start" timestamp(0) DEFAULT now() NOT NULL
)
WITH (OIDS=FALSE)

;

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS "public"."users";
CREATE TABLE "public"."users" (
"created_at" timestamp(6) DEFAULT now() NOT NULL,
"updated_at" timestamp(6) DEFAULT now() NOT NULL,
"name" varchar(20) COLLATE "default" NOT NULL,
"email" varchar(40) COLLATE "default" NOT NULL,
"password" varchar(100) COLLATE "default" NOT NULL,
"gold" int4 DEFAULT 1000 NOT NULL,
"diamonds" int4 DEFAULT 0 NOT NULL,
"darkwood" int4 DEFAULT 0 NOT NULL,
"runestone" int4 DEFAULT 0 NOT NULL,
"veritium" int4 DEFAULT 0 NOT NULL,
"trueseed" int4 DEFAULT 0 NOT NULL,
"rank" int2 DEFAULT 0 NOT NULL,
"alliance_name" varchar(20) COLLATE "default",
"alliance_rank" varchar(15) COLLATE "default"
)
WITH (OIDS=FALSE)

;

-- ----------------------------
-- Alter Sequences Owned By 
-- ----------------------------

-- ----------------------------
-- Uniques structure for table alliances
-- ----------------------------
ALTER TABLE "public"."alliances" ADD UNIQUE ("name");

-- ----------------------------
-- Indexes structure for table cities
-- ----------------------------
CREATE UNIQUE INDEX "city_coord_index" ON "public"."cities" USING btree ("x", "y", "continent_x", "continent_y");
ALTER TABLE "public"."cities" CLUSTER ON "city_coord_index";
CREATE INDEX "city_username_index" ON "public"."cities" USING btree ("user_name");

-- ----------------------------
-- Uniques structure for table cities
-- ----------------------------
ALTER TABLE "public"."cities" ADD UNIQUE ("x", "y");

-- ----------------------------
-- Primary Key structure for table cities
-- ----------------------------
ALTER TABLE "public"."cities" ADD PRIMARY KEY ("x", "y", "continent_x", "continent_y");

-- ----------------------------
-- Indexes structure for table constructions
-- ----------------------------
CREATE UNIQUE INDEX "city_index" ON "public"."constructions" USING btree ("x", "y", "city_x", "city_y");
ALTER TABLE "public"."constructions" CLUSTER ON "city_index";

-- ----------------------------
-- Uniques structure for table constructions
-- ----------------------------
ALTER TABLE "public"."constructions" ADD UNIQUE ("x", "y", "city_x", "city_y");

-- ----------------------------
-- Primary Key structure for table constructions
-- ----------------------------
ALTER TABLE "public"."constructions" ADD PRIMARY KEY ("city_x", "x", "y", "city_y");

-- ----------------------------
-- Indexes structure for table continents
-- ----------------------------
CREATE UNIQUE INDEX "xy_index" ON "public"."continents" USING btree ("x", "y");
ALTER TABLE "public"."continents" CLUSTER ON "xy_index";

-- ----------------------------
-- Uniques structure for table continents
-- ----------------------------
ALTER TABLE "public"."continents" ADD UNIQUE ("x", "y");

-- ----------------------------
-- Primary Key structure for table continents
-- ----------------------------
ALTER TABLE "public"."continents" ADD PRIMARY KEY ("x", "y");

-- ----------------------------
-- Indexes structure for table dungeons
-- ----------------------------
CREATE UNIQUE INDEX "coord_dungeon_index" ON "public"."dungeons" USING btree ("continent_x", "continent_y", "x", "y");
ALTER TABLE "public"."dungeons" CLUSTER ON "coord_dungeon_index";

-- ----------------------------
-- Uniques structure for table dungeons
-- ----------------------------
ALTER TABLE "public"."dungeons" ADD UNIQUE ("x", "y");

-- ----------------------------
-- Primary Key structure for table dungeons
-- ----------------------------
ALTER TABLE "public"."dungeons" ADD PRIMARY KEY ("continent_x", "continent_y", "x", "y");

-- ----------------------------
-- Indexes structure for table military_actions
-- ----------------------------
CREATE INDEX "origin_index" ON "public"."military_actions" USING btree ("origin_id");
CREATE INDEX "target_index" ON "public"."military_actions" USING btree ("target_id");
CREATE INDEX "arrival_index" ON "public"."military_actions" USING btree ("arrival");

-- ----------------------------
-- Primary Key structure for table military_actions
-- ----------------------------
ALTER TABLE "public"."military_actions" ADD PRIMARY KEY ("id");

-- ----------------------------
-- Indexes structure for table tiles
-- ----------------------------
CREATE UNIQUE INDEX "coord_tile_index" ON "public"."tiles" USING btree ("x", "y", "continent_x", "continent_y");
ALTER TABLE "public"."tiles" CLUSTER ON "coord_tile_index";

-- ----------------------------
-- Uniques structure for table tiles
-- ----------------------------
ALTER TABLE "public"."tiles" ADD UNIQUE ("x", "y");

-- ----------------------------
-- Primary Key structure for table tiles
-- ----------------------------
ALTER TABLE "public"."tiles" ADD PRIMARY KEY ("x", "y", "continent_x", "continent_y");

-- ----------------------------
-- Indexes structure for table upgrades
-- ----------------------------
CREATE INDEX "order_index" ON "public"."upgrades" USING btree ("index");
CREATE UNIQUE INDEX "tile_upgrade_index" ON "public"."upgrades" USING btree ("x", "y", "city_x", "city_y", "index");
ALTER TABLE "public"."upgrades" CLUSTER ON "tile_upgrade_index";

-- ----------------------------
-- Uniques structure for table upgrades
-- ----------------------------
ALTER TABLE "public"."upgrades" ADD UNIQUE ("x", "y", "city_x", "city_y", "index");

-- ----------------------------
-- Primary Key structure for table upgrades
-- ----------------------------
ALTER TABLE "public"."upgrades" ADD PRIMARY KEY ("x", "y", "city_x", "city_y");

-- ----------------------------
-- Indexes structure for table users
-- ----------------------------
CREATE UNIQUE INDEX "email_index" ON "public"."users" USING btree ("email");
CREATE UNIQUE INDEX "user_name_index" ON "public"."users" USING btree ("name");

-- ----------------------------
-- Uniques structure for table users
-- ----------------------------
ALTER TABLE "public"."users" ADD UNIQUE ("name");
ALTER TABLE "public"."users" ADD UNIQUE ("email");

-- ----------------------------
-- Primary Key structure for table users
-- ----------------------------
ALTER TABLE "public"."users" ADD PRIMARY KEY ("email", "name");

-- ----------------------------
-- Foreign Key structure for table "public"."cities"
-- ----------------------------
ALTER TABLE "public"."cities" ADD FOREIGN KEY ("x", "y", "continent_x", "continent_y") REFERENCES "public"."tiles" ("x", "y", "continent_x", "continent_y") ON DELETE CASCADE ON UPDATE CASCADE;
ALTER TABLE "public"."cities" ADD FOREIGN KEY ("user_name") REFERENCES "public"."users" ("name") ON DELETE SET NULL ON UPDATE CASCADE;

-- ----------------------------
-- Foreign Key structure for table "public"."constructions"
-- ----------------------------
ALTER TABLE "public"."constructions" ADD FOREIGN KEY ("city_x", "city_y") REFERENCES "public"."cities" ("x", "y") ON DELETE CASCADE ON UPDATE CASCADE;

-- ----------------------------
-- Foreign Key structure for table "public"."dungeons"
-- ----------------------------
ALTER TABLE "public"."dungeons" ADD FOREIGN KEY ("x", "y", "continent_x", "continent_y") REFERENCES "public"."tiles" ("x", "y", "continent_x", "continent_y") ON DELETE CASCADE ON UPDATE CASCADE;

-- ----------------------------
-- Foreign Key structure for table "public"."tiles"
-- ----------------------------
ALTER TABLE "public"."tiles" ADD FOREIGN KEY ("continent_x", "continent_y") REFERENCES "public"."continents" ("x", "y") ON DELETE CASCADE ON UPDATE CASCADE;

-- ----------------------------
-- Foreign Key structure for table "public"."upgrades"
-- ----------------------------
ALTER TABLE "public"."upgrades" ADD FOREIGN KEY ("x", "y", "city_x", "city_y") REFERENCES "public"."constructions" ("x", "y", "city_x", "city_y") ON DELETE CASCADE ON UPDATE CASCADE;

-- ----------------------------
-- Foreign Key structure for table "public"."users"
-- ----------------------------
ALTER TABLE "public"."users" ADD FOREIGN KEY ("alliance_name") REFERENCES "public"."alliances" ("name") ON DELETE SET NULL ON UPDATE CASCADE;
