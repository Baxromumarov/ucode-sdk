CREATE TABLE "Справочник.measurements" (
  "guid" uuid,
  "base_quantity" varchar,
  "base_unit" varchar,
  "symbol" varchar
);

CREATE TABLE "country_parent" (
  "guid" uuid PRIMARY KEY, 
  "name" varchar
);

CREATE TABLE "countries" (
  "guid" uuid PRIMARY KEY,
  "name" varchar,
  "code" varchar,
  "flag" varchar,
  "country_parent_id.name" uuid
);

CREATE TABLE "time_zones" (
  "guid" uuid PRIMARY KEY,
  "name" varchar,
  "offset" int
);

CREATE TABLE "cities" (
  "guid" uuid PRIMARY KEY,
  "name" varchar,
  "code" varchar,
  "country_id.name" uuid,
  "location" point,
  "geofence" Polygon,
  "time_zone_id.name" uuid
);

CREATE TABLE "currencies" (
  "guid" uuid PRIMARY KEY,
  "name" varchar,
  "short_name" varchar,
  "code" varchar,
  "country_id.name" uuid,
  "photo" varchar
);

CREATE TABLE "districts" (
  "guid" uuid PRIMARY KEY,
  "name" varchar,
  "city_id.name" uuid,
  "location" point,
  "geofence" Polygon
);

ALTER TABLE
  "country"
ADD
  FOREIGN KEY ("country_parent_id") REFERENCES "country_parent" ("guid");

ALTER TABLE
  "currency"
ADD
  FOREIGN KEY ("country_id") REFERENCES "country" ("guid");

ALTER TABLE
  "city"
ADD
  FOREIGN KEY ("time_zone_id") REFERENCES "time_zone" ("guid");

ALTER TABLE
  "city"
ADD
  FOREIGN KEY ("country_id") REFERENCES "country" ("guid");

ALTER TABLE
  "district"
ADD
  FOREIGN KEY ("city_id") REFERENCES "city" ("guid");