CREATE TABLE "rr_modul.rrr_table" (
  "guid" uuid PRIMARY KEY,
  "balance" varchar
);

CREATE TABLE "rrrr_table" (
  "guid" uuid PRIMARY KEY,
  "name" varchar
);

CREATE TABLE "rrrrr_product" (
  "guid" uuid PRIMARY KEY,
  "product_name" varchar,
  "rrr_table.balance" uuid,
  "rrrr_table.name" uuid
);


ALTER TABLE "sdk_product" ADD FOREIGN KEY ("sdk_table_id") REFERENCES "sdk_table" ("guid");
