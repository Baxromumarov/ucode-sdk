CREATE TABLE "sdk-module.sdk_table" (
  "guid" uuid PRIMARY KEY,
  "balance" float
);

CREATE TABLE "sdk_product" (
  "guid" uuid PRIMARY KEY,
  "product_name" varchar,
  "sdk_table_id.balance" uuid
);


ALTER TABLE "sdk_product" ADD FOREIGN KEY ("sdk_table_id") REFERENCES "sdk_table" ("guid");
